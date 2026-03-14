package services

import (
	"errors"
	"hoodhire/internal/repositories"
	"hoodhire/structures/dto"
	"hoodhire/structures/models"
)

type TicketServices struct {
	Repo *repositories.TicketRepo
}

func NewTicketServices(r *repositories.TicketRepo) *TicketServices {
	return &TicketServices{Repo: r}
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Helper~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *TicketServices) getSeekerByUserID(userID uint) (*models.Seeker, error) {
	var seeker models.Seeker
	err := s.Repo.DB.Where("user_id = ?", userID).First(&seeker).Error
	if err != nil {
		return nil, errors.New("seeker profile not found")
	}
	return &seeker, nil
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Seeker~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *TicketServices) CreateTicket(userID uint, input *dto.CreateTicketDTO) error {
	seeker, err := s.getSeekerByUserID(userID)
	if err != nil {
		return err
	}
	ticket := &models.Ticket{
		SeekerID:    seeker.ID,
		BusinessID:  input.BusinessID,
		Type:        input.Type,
		Subject:     input.Subject,
		Description: input.Description,
	}
	return s.Repo.CreateTicket(ticket)
}

func (s *TicketServices) GetMyTickets(userID uint) ([]models.Ticket, error) {
	seeker, err := s.getSeekerByUserID(userID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetMyTickets(seeker.ID)
}

func (s *TicketServices) DeleteTicket(userID, ticketID uint) error {
	seeker, err := s.getSeekerByUserID(userID)
	if err != nil {
		return err
	}
	return s.Repo.DeleteTicket(ticketID, seeker.ID)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Admin~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *TicketServices) GetAllTickets() ([]models.Ticket, error) {
	return s.Repo.GetAllTickets()
}

func (s *TicketServices) GetTicketsByType(ticketType string) ([]models.Ticket, error) {
	return s.Repo.GetTicketsByType(ticketType)
}

func (s *TicketServices) GetTicketsByStatus(status string) ([]models.Ticket, error) {
	return s.Repo.GetTicketsByStatus(status)
}

func (s *TicketServices) UpdateTicketStatus(ticketID uint, input *dto.UpdateTicketStatusDTO) error {
	return s.Repo.UpdateTicketStatus(ticketID, input.Status)
}

func (s *TicketServices) GetTicketsByBusiness(businessID uint) ([]models.Ticket, error) {
	return s.Repo.GetTicketsByBusiness(businessID)
}