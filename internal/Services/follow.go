package services

import (
	"errors"
	"hoodhire/internal/repositories"
	"hoodhire/structures/dto"
	"hoodhire/structures/models"
)

type FollowServices struct {
	Repo *repositories.FollowRepo
}

func NewFollowService(follorepo *repositories.FollowRepo)*FollowServices{
	return &FollowServices{Repo: follorepo}
}

func (s *FollowServices)GetSeekerByuser(userID uint)(*models.Seeker,error){
	var seeker models.Seeker
	err:= s.Repo.DB.Where("user_id = ?",userID).First(&seeker).Error
	if err!=nil{
		return nil,err
	}
	return &seeker,nil
}

func (s *FollowServices)FollowBusiness(userID,businessID uint)error{
	seeker,err:=s.GetSeekerByuser(userID)
	if err!=nil{
		return err
	}
	if s.Repo.IsFollowing(seeker.ID,businessID){
		return errors.New("already following")
	}
	return s.Repo.FollowBusiness(seeker.ID,businessID)
}

func (s *FollowServices)UnfollowBusiness(userID,businessID uint)error{
	seeker,err:=s.GetSeekerByuser(userID)
	if err!=nil{
		return err
	}
	if !s.Repo.IsFollowing(seeker.ID,businessID){
		return errors.New("not followed")
	}
	return s.Repo.UnfollowBusinsess(seeker.ID,businessID)
}

func (s *FollowServices) GetFollowedBusinesses(userID uint) ([]models.BusinessFollow, error) {
	seeker, err := s.GetSeekerByuser(userID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetFollows(seeker.ID)

}

func (s *FollowServices) IsFollowing(userID, businessID uint) (bool, error) {
    seeker, err := s.GetSeekerByuser(userID)
    if err != nil {
        return false, err
    }
    return s.Repo.IsFollowing(seeker.ID, businessID), nil
}


//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Helper~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *FollowServices) getSeekerByUserID(userID uint) (*models.Seeker, error) {
	var seeker models.Seeker
	err := s.Repo.DB.Where("user_id = ?", userID).First(&seeker).Error
	if err != nil {
		return nil, errors.New("seeker profile not found")
	}
	return &seeker, nil
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Create Review~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *FollowServices) CreateReview(userID, businessID uint, input *dto.CreateReviewDto) error {
	seeker, err := s.getSeekerByUserID(userID)
	if err != nil {
		return err
	}
	if s.Repo.HasReviewed(seeker.ID, businessID) {
		return errors.New("you have already reviewed this business")
	}
	review := &models.BusinessReview{
		SeekerID:   seeker.ID,
		BusinessID: businessID,
		Rating:     input.Rating,
		Message:    input.Message,
	}
	return s.Repo.CreateReview(review)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Update Review~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *FollowServices) UpdateReview(userID, businessID uint, input *dto.CreateReviewDto) error {
	seeker, err := s.getSeekerByUserID(userID)
	if err != nil {
		return err
	}
	if !s.Repo.HasReviewed(seeker.ID, businessID) {
		return errors.New("you have not reviewed this business yet")
	}
	return s.Repo.UpdateReview(seeker.ID, businessID, input.Rating, input.Message)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Delete Review~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *FollowServices) DeleteReview(userID, businessID uint) error {
	seeker, err := s.getSeekerByUserID(userID)
	if err != nil {
		return err
	}
	return s.Repo.DeleteReview(seeker.ID, businessID)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Get Reviews~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *FollowServices) GetReviewsByBusiness(businessID uint) ([]models.BusinessReview, error) {
	return s.Repo.GetReviewsByBusiness(businessID)
}

func (s *FollowServices) GetMyReview(userID, businessID uint) (*models.BusinessReview, error) {
	seeker, err := s.getSeekerByUserID(userID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetMyReview(seeker.ID, businessID)
}