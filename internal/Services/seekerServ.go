package services

import (
	"errors"
	repositories "hoodhire/internal/repositories"
	dto "hoodhire/structures/dto"
	"hoodhire/structures/models"
)

type SeekerServices struct {
	Repo *repositories.SeekerRepo
}

func NewSeekerServices(r *repositories.SeekerRepo) *SeekerServices {
	return &SeekerServices{Repo: r}
}


//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Create Seeker~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *SeekerServices) CreateSeeker(userID uint, input *dto.CreateSeekerDTO) (bool, error) {
	if s.Repo.SeekerExist(userID) {
		return false, errors.New("seeker profile already exists")
	}

	seeker := &models.Seeker{
		UserID:         userID,
		FullName:       input.FullName,
		Age:            input.Age,
		PhoneNumber:    input.PhoneNumber,
		Gender:         input.Gender,
		CurrentStatus:  input.CurrentStatus,
		Bio:            input.Bio,
		About: input.About,
		CurrentAddress: input.CurrentAddress,
		Locality:       input.Locality,
		IsCompleted:    true, 
	}

	edu := &models.Education{
		FieldOfStudy:   input.FieldOfStudy,
		CourseName:     input.CourseName,
		InstituteName:  input.InstituteName,
		StartYear:      input.StartYear,
		GraduationYear: input.GraduationYear,
		IsOngoing:      input.IsOngoing,
	}

	return true, s.Repo.CreateSeekerWithEducation(seeker, edu,input.CategoryIDs)
}


//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Fetch seeker profile ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func (s *SeekerServices) GetSeeker(userID uint) (*models.Seeker, error) {
	return s.Repo.GetSeeker(userID)
}
func (s *SeekerServices)GetSeekerByID(seekerID uint)(*models.Seeker,error){
	return s.Repo.GetSeekerByID(seekerID)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Update seeker profile ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *SeekerServices) UpdateSeeker(userID uint, input *dto.CreateSeekerDTO) (*models.Seeker, error) {
	seeker, err := s.GetSeeker(userID)
	if err != nil {
		return nil, err
	}
	
	seeker.FullName = input.FullName
	seeker.Age = input.Age
	seeker.PhoneNumber = input.PhoneNumber
	seeker.Gender = input.Gender
	seeker.CurrentStatus = input.CurrentStatus
	seeker.Bio = input.Bio
	seeker.About=input.About
	seeker.CurrentAddress = input.CurrentAddress
	seeker.Locality = input.Locality
	seeker.IsCompleted = true
	
	edu := &models.Education{
		SeekerID:       seeker.ID,
		FieldOfStudy:   input.FieldOfStudy,
		CourseName:     input.CourseName,
		InstituteName:  input.InstituteName,
		StartYear:      input.StartYear,
		GraduationYear: input.GraduationYear,
		IsOngoing:      input.IsOngoing,
	}
	
	return seeker, s.Repo.UpdateSeekerWithEducation(seeker, edu, input.CategoryIDs) 
}


func (s *SeekerServices) DeleteSeeker(userID uint) error {
	return s.Repo.DeleteSeeker(userID)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Education ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func (s *SeekerServices) UpsertEducation(userID uint, input *dto.UpdateEducationDTO) error {
	seeker, err := s.GetSeeker(userID)
	if err != nil {
		return err
	}
	edu := &models.Education{
		SeekerID:       seeker.ID,
		FieldOfStudy:   input.FieldOfStudy,
		CourseName:     input.CourseName,
		InstituteName:  input.InstituteName,
		StartYear:      input.StartYear,
		GraduationYear: input.GraduationYear,
		IsOngoing:      input.IsOngoing,
	}
	return s.Repo.UpsertEducation(edu)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Work Experience~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func (s *SeekerServices) AddWorkExperience(userID uint, input *dto.WorkExperienceDTO) error {
	seeker, err := s.GetSeeker(userID)
	if err != nil {
		return err
	}
	exp := &models.WorkExperience{
		SeekerID:     seeker.ID,
		CompanyName:  input.CompanyName,
		Position:     input.Position,
		Duration:     input.Duration,
		IsCurrentJob: input.IsCurrentJob,
		Description:  input.Description,
	}
	return s.Repo.AddWorkExperience(exp)
}

func (s *SeekerServices) GetWorkExperiences(userID uint) ([]models.WorkExperience, error) {
	seeker, err := s.GetSeeker(userID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetWorkExperiences(seeker.ID)
}

func (s *SeekerServices) DeleteWorkExperience(userID uint, expID uint) error {
	seeker, err := s.GetSeeker(userID)
	if err != nil {
		return err
	}
	return s.Repo.DeleteWorkExperience(expID, seeker.ID)
}

func IsComplete(s *models.Seeker, edu *models.Education) bool {
	basicInfoFilled := s.FullName != "" &&
	s.Gender != "" &&
	s.Age > 0 &&
	s.PhoneNumber != "" &&
	s.CurrentStatus != "" &&
	s.CurrentAddress != "" &&
	s.Locality != ""
	
	educationFilled := edu != nil &&
	edu.CourseName != "" &&
	edu.InstituteName != "" &&
	edu.StartYear > 0
	
	return basicInfoFilled && educationFilled
}



//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Work Preference~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func (s *SeekerServices) UpsertWorkPreference(userID uint, input *dto.WorkPreferenceDTO) error {
	seeker, err := s.GetSeeker(userID)
	if err != nil {
		return err
	}
	pref := &models.WorkPreference{
		SeekerID:       seeker.ID,
		Monday:         input.Monday,
		Tuesday:        input.Tuesday,
		Wednesday:      input.Wednesday,
		Thursday:       input.Thursday,
		Friday:         input.Friday,
		Saturday:       input.Saturday,
		Sunday:         input.Sunday,
		PreferredShift: input.PreferredShift,
		PartTime:       input.PartTime,
		FullTime:       input.FullTime,
		Immediate:      input.Immediate,
	}
	return s.Repo.UpsertWorkPreference(pref)
}

func (s *SeekerServices) GetWorkPreference(userID uint) (*models.WorkPreference, error) {
	seeker, err := s.GetSeeker(userID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetWorkPreference(seeker.ID)
	}													
	
	
//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Job Interests~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *SeekerServices) UpdateJobInterests(userID uint, input *dto.JobInterestsDTO) error {
	seeker, err := s.GetSeeker(userID)
	if err != nil {
		return err
	}
	return s.Repo.UpdateJobIntereset(seeker.ID, input.CategoryIDs)
}

func (s *SeekerServices) GetJobCategories() ([]models.JobCategory, error) {
	return s.Repo.GetJobCategories()
}