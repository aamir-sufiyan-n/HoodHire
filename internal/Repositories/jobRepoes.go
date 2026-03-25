package repositories

import (
	"hoodhire/structures/models"
	"time"

	"gorm.io/gorm"
)

type JobRepo struct {
	DB *gorm.DB
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Create Jobs~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (r *JobRepo) CreateJob(job *models.Job, desc *models.JobDescription) error {

	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(job).Error; err != nil {
			return err
		}
		desc.JobID = job.ID
		if err := tx.Create(desc).Error; err != nil {
			return err
		}
		return nil
	})
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Get Jobs~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (r *JobRepo) GetJobsByID(id uint) (*models.Job, error) {
	var job models.Job
	err := r.DB.Joins("Description").Joins("Business").Joins("Category").
		Where("jobs.id=?", id).First(&job).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *JobRepo) GetJobsByHirer(hirerID uint) ([]models.Job, error) {
	var jobs []models.Job
	err := r.DB.Preload("Description").
		Preload("Business").
		Preload("Category").
		Where("hirer_id = ?", hirerID).Find(&jobs).Error
	return jobs, err
}

func (r *JobRepo) GetActiveJobs() ([]models.Job, error) {
	var jobs []models.Job
	err := r.DB.
		Joins("Description").
		Joins("Business").
		Joins("Category").
		Where("jobs.status = ?", "open").
		Find(&jobs).Error
	return jobs, err
}

func (r *JobRepo) GetJobsByCategory(categoryID uint) ([]models.Job, error) {
	var jobs []models.Job
	err := r.DB.Preload("Description").
		Preload("Business").
		Preload("Category").
		Where("jobs.category_id = ?  AND jobs.status = ?", categoryID, "open").
		Find(&jobs).Error
	return jobs, err
}

func (r *JobRepo) GetJobsByLocality(locality string) ([]models.Job, error) {
	var jobs []models.Job
	err := r.DB.Preload("Description").
		Preload("Business").
		Preload("Category").
		Joins("JOIN businesses ON businesses.id = jobs.business_id").
		Where("businesses.locality = ? AND jobs.status = ?", locality,"open").
		Find(&jobs).Error
	return jobs, err
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Update Job~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (r *JobRepo) UpdateJobWithDescription(job *models.Job, desc *models.JobDescription) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(job).Error; err != nil {
			return err
		}
		var existing models.JobDescription
		err := tx.Where("job_id = ?", job.ID).First(&existing).Error
		if err != nil {
			return err
		}
		desc.ID = existing.ID
		desc.JobID = job.ID
		return tx.Model(desc).Select("*").Save(desc).Error 
	})
}

func (r *JobRepo) UpdateJobStatus(jobID uint, status string) error {
	return r.DB.Model(&models.Job{}).
		Where("id = ?", jobID).
		Updates(map[string]interface{}{
			"status":    status,
		}).Error
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Delete Jobs~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (r *JobRepo) DeleteJob(jobID, hirerID uint) error {
	return r.DB.Where("id=? AND hirer_id=?", jobID, hirerID).Delete(&models.Job{}).Error
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Belongs to Hirer~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (r *JobRepo) JobBelongsToHirer(jobID uint, hirerID uint) bool {
	err := r.DB.Where("id = ? AND hirer_id = ?", jobID, hirerID).First(&models.Job{}).Error
	return err == nil
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Applications~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (r *JobRepo) ApplytoJob(application *models.JobApplication) error {
	return r.DB.Create(application).Error
}

func (r *JobRepo) AlreadyApplied(jobID, seekerID uint) bool {
	err := r.DB.Where("job_id=? AND seeker_id =?", jobID, seekerID).First(&models.JobApplication{}).Error
	return err == nil
}

func (r *JobRepo) GetApplicationsByJob(jobID uint) ([]models.JobApplication, error) {
	var applications []models.JobApplication
	err := r.DB.
		Preload("Seeker.User").
		Preload("Seeker.Education").
		Preload("Seeker.WorkExperiences").
		Preload("Seeker.WorkPreference").
		Preload("Seeker.JobInterests.Category").
		Where("job_id = ?", jobID).Find(&applications).Error
	if err != nil {
		return nil, err
	}
	for i := range applications {
		applications[i].Seeker.Email = applications[i].Seeker.User.Email
	}
	return applications, nil
}
func (r *JobRepo) GetApplicationsBySeeker(seekerID uint) ([]models.JobApplication, error) {
	var applications []models.JobApplication
	err := r.DB.Preload("Job").
		Preload("Job.Description").
		Preload("Job.Business").
		Where("seeker_id = ?", seekerID).Find(&applications).Error
	return applications, err
}

func (r *JobRepo) UpdateApplicationStatus(applicationID uint, status string) error {
	return r.DB.Model(&models.JobApplication{}).
		Where("id = ?", applicationID).
		Update("status", status).Error
}

func (r *JobRepo) WithdrawApplication(applicationID uint, seekerID uint) error {
	return r.DB.Where("id = ? AND seeker_id = ?", applicationID, seekerID).
		Delete(&models.JobApplication{}).Error
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Deadline Cleanup~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (r *JobRepo) CloseExpiredJobs() error {
	return r.DB.Model(&models.Job{}).
		Where("deadline < ? AND status = ?", time.Now(), "open").
		Updates(map[string]interface{}{
			"status":    "closed",
		}).Error
}



