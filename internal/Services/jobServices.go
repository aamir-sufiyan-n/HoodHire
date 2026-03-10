package services

import (
	"errors"
	repositories "hoodhire/internal/repositories"
	dto "hoodhire/structures/dto"
	"hoodhire/structures/models"
	"time"
)

type JobServices struct {
	Repo      *repositories.JobRepo
	HirerRepo *repositories.HirerRepo
}

func NewJobServices(r *repositories.JobRepo, h *repositories.HirerRepo) *JobServices {
	return &JobServices{Repo: r, HirerRepo: h}
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Create Job~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *JobServices) CreateJob(userID uint, input *dto.CreateJobDTO) error {
	hirer, err := s.HirerRepo.GetHirer(userID)
	if err != nil {
		return errors.New("hirer profile not found")
	}
	if hirer.Business == nil {
		return errors.New("business profile not found")
	}
	if !hirer.Business.IsVerified {
		return errors.New("business is not verified yet, cannot post jobs")
	}

	job := &models.Job{
		HirerID:    hirer.ID,
		BusinessID: hirer.Business.ID,
		CategoryID: input.CategoryID,
		Status:     "open",
		Deadline:   input.Deadline,
	}

	desc := &models.JobDescription{
		Title:               input.Title,
		Description:         input.Description,
		JobType:             input.JobType,
		Shift:               input.Shift,
		Duration:            input.Duration,
		SalaryMin:           input.SalaryMin,
		SalaryMax:           input.SalaryMax,
		SalaryType:          input.SalaryType,
		MinAge:              input.MinAge,
		MaxAge:              input.MaxAge,
		GenderPref:          input.GenderPref,
		ExperienceRequired:  input.ExperienceRequired,
		Monday:              input.Monday,
		Tuesday:             input.Tuesday,
		Wednesday:           input.Wednesday,
		Thursday:            input.Thursday,
		Friday:              input.Friday,
		Saturday:            input.Saturday,
		Sunday:              input.Sunday,
		KeyResponsibilities: input.KeyResponsibilities, 
		Skills:              input.Skills,
	}

	return s.Repo.CreateJob(job, desc)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Fetch Jobs~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *JobServices) GetJobByID(jobID uint) (*models.Job, error) {
	return s.Repo.GetJobsByID(jobID)
}

func (s *JobServices) GetMyJobs(userID uint) ([]models.Job, error) {
	hirer, err := s.HirerRepo.GetHirer(userID)
	if err != nil {
		return nil, errors.New("hirer profile not found")
	}
	return s.Repo.GetJobsByHirer(hirer.ID)
}

func (s *JobServices) GetActiveJobs() ([]models.Job, error) {
	return s.Repo.GetActiveJobs()
}

func (s *JobServices) GetJobsByCategory(categoryID uint) ([]models.Job, error) {
	return s.Repo.GetJobsByCategory(categoryID)
}

func (s *JobServices) GetJobsByLocality(locality string) ([]models.Job, error) {
	return s.Repo.GetJobsByLocality(locality)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Update Job~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *JobServices) UpdateJob(userID uint, jobID uint, input *dto.UpdateJobDTO) error {
	hirer, err := s.HirerRepo.GetHirer(userID)
	if err != nil {
		return errors.New("hirer profile not found")
	}
	if !s.Repo.JobBelongsToHirer(jobID, hirer.ID) {
		return errors.New("unauthorized")
	}

	job, err := s.Repo.GetJobsByID(jobID)
	if err != nil {
		return err
	}

	if input.CategoryID != 0 {
		job.CategoryID = input.CategoryID
	}
	if input.Status != "" {
		job.Status = input.Status
	}
	if input.Deadline != nil {
		job.Deadline = input.Deadline
	}

	desc := &models.JobDescription{
		Title:              input.Title,
		Description:        input.Description,
		JobType:            input.JobType,
		Shift:              input.Shift,
		Duration:           input.Duration,
		SalaryMin:          input.SalaryMin,
		SalaryMax:          input.SalaryMax,
		SalaryType:         input.SalaryType,
		MinAge:             input.MinAge,
		MaxAge:             input.MaxAge,
		GenderPref:         input.GenderPref,
		ExperienceRequired: input.ExperienceRequired,
		Monday:             input.Monday,
		Tuesday:            input.Tuesday,
		Wednesday:          input.Wednesday,
		Thursday:           input.Thursday,
		Friday:             input.Friday,
		Saturday:           input.Saturday,
		Sunday:             input.Sunday,
		KeyResponsibilities: input.KeyResponsibilities,
		Skills: input.Skills,
	}

	return s.Repo.UpdateJobWithDescription(job, desc)
}

func (s *JobServices) UpdateJobStatus(userID uint, jobID uint, input *dto.UpdateJobStatusDTO) error {
	hirer, err := s.HirerRepo.GetHirer(userID)
	if err != nil {
		return errors.New("hirer profile not found")
	}
	if !s.Repo.JobBelongsToHirer(jobID, hirer.ID) {
		return errors.New("unauthorized")
	}
	return s.Repo.UpdateJobStatus(jobID, input.Status)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Delete Job~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *JobServices) DeleteJob(userID uint, jobID uint) error {
	hirer, err := s.HirerRepo.GetHirer(userID)
	if err != nil {
		return errors.New("hirer profile not found")
	}
	if !s.Repo.JobBelongsToHirer(jobID, hirer.ID) {
		return errors.New("unauthorized")
	}
	return s.Repo.DeleteJob(jobID, hirer.ID)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Applications~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *JobServices) ApplyToJob(userID uint, jobID uint, input *dto.JobApplicationDTO) error {
	seeker, err := s.getSeekerByUserID(userID)
	if err != nil {
		return err
	}

	job, err := s.Repo.GetJobsByID(jobID)
	if err != nil {
		return errors.New("job not found")
	}
	if job.Status != "open" {
		return errors.New("job is no longer accepting applications")
	}
	if job.Deadline != nil && job.Deadline.Before(time.Now()) {
		return errors.New("application deadline has passed")
	}
	if s.Repo.AlreadyApplied(jobID, seeker.ID) {
		return errors.New("you have already applied to this job")
	}

	application := &models.JobApplication{
		JobID:    jobID,
		SeekerID: seeker.ID,
		Status:   "pending",
		Message:  input.Message,
	}
	return s.Repo.ApplytoJob(application)
}

func (s *JobServices) GetApplicationsForJob(userID uint, jobID uint) ([]models.JobApplication, error) {
	hirer, err := s.HirerRepo.GetHirer(userID)
	if err != nil {
		return nil, errors.New("hirer profile not found")
	}
	if !s.Repo.JobBelongsToHirer(jobID, hirer.ID) {
		return nil, errors.New("unauthorized")
	}
	return s.Repo.GetApplicationsByJob(jobID)
}

func (s *JobServices) GetMyApplications(userID uint) ([]models.JobApplication, error) {
	seeker, err := s.getSeekerByUserID(userID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetApplicationsBySeeker(seeker.ID)
}

func (s *JobServices) UpdateApplicationStatus(userID uint, applicationID uint, input *dto.UpdateApplicationStatusDTO) error {
	hirer, err := s.HirerRepo.GetHirer(userID)
	if err != nil {
		return errors.New("hirer profile not found")
	}
	// verify the application belongs to one of this hirer's jobs
	var application models.JobApplication
	if err := s.Repo.DB.First(&application, applicationID).Error; err != nil {
		return errors.New("application not found")
	}
	if !s.Repo.JobBelongsToHirer(application.JobID, hirer.ID) {
		return errors.New("unauthorized")
	}
	return s.Repo.UpdateApplicationStatus(applicationID, input.Status)
}

func (s *JobServices) WithdrawApplication(userID uint, applicationID uint) error {
	seeker, err := s.getSeekerByUserID(userID)
	if err != nil {
		return err
	}
	return s.Repo.WithdrawApplication(applicationID, seeker.ID)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Helper~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// avoids importing SeekerRepo — fetches seeker directly via DB on JobRepo
func (s *JobServices) getSeekerByUserID(userID uint) (*models.Seeker, error) {
	var seeker models.Seeker
	err := s.Repo.DB.Where("user_id = ?", userID).First(&seeker).Error
	if err != nil {
		return nil, errors.New("seeker profile not found")
	}
	return &seeker, nil
}
