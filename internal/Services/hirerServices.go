package services

import (
	"errors"
	"hoodhire/internal/repositories"
	"hoodhire/structures/dto"
	"hoodhire/structures/models"
)

type HirerServices struct {
	Repo *repositories.HirerRepo
}

func NewHirerService(r *repositories.HirerRepo)*HirerServices{
	return &HirerServices{Repo:r}
}

func (s *HirerServices)CreateHirer(userID uint,input *dto.CreateHirerDto)error{
	if s.Repo.HirerExists(userID){
		return errors.New("user already exists")
	}
	hirer:=&models.Hirer{
		UserID: userID,
		FullName: input.FullName,
		PhoneNumber: input.PhoneNumber,
		CurrentAddress: input.CurrentAddress,
		IsProfileComplete: true,
	}
	if err:= s.Repo.Create(hirer); err!=nil{
		return err
	}
	business:=&models.Business{
		HirerID: hirer.ID,
		BusinessName: input.BusinessName,
		BusinessPhone: input.BusinessPhone,
		Niche: input.Niche,
		Address: input.Address,
		Locality: input.Locality,
		Bio: input.Bio,
	}
	
	return s.Repo.CreateBusiness(business)
}

func ( s *HirerServices)GetHirer(userid uint)(*models.Hirer,error){
	hirer,err:=s.Repo.GetHirer(userid)
	if err!=nil{
		return nil,err
	}
	if hirer==nil{
		return nil,errors.New("hirer profile not found")
	}
	return hirer,nil
}

func (s *HirerServices) UpdateHirer(userid uint, input *dto.CreateHirerDto) (*models.Hirer, error) {
	hirer, err := s.Repo.GetHirer(userid)
	if err != nil {
		return nil, err
	}

	// Update personal info
	hirer.FullName = input.FullName
	hirer.CurrentAddress = input.CurrentAddress
	hirer.PhoneNumber = input.PhoneNumber

	if err := s.Repo.Update(hirer); err != nil {
		return nil, err
	}

	// Update business info
	business, err := s.Repo.GetBusiness(hirer.ID)
	if err != nil {
		return nil, err
	}
	
	business.BusinessName = input.BusinessName
	business.Address = input.Address
	business.Niche=input.Niche
	business.BusinessPhone = input.BusinessPhone
	business.Locality = input.Locality
	business.Bio = input.Bio

	if err := s.Repo.UpdateBusiness(business); err != nil {
		return nil, err
	}

	return hirer, nil
}
func (s *HirerServices)DeleteHirer( userid uint)error{
	hirer,err:=s.GetHirer(userid)
	if err!=nil{
		return err
	}
	return s.Repo.Delete(hirer)
}
