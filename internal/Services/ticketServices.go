package services

import (
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

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Seeker & Hirer~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *TicketServices) CreateTicket(userID uint, role string, input *dto.CreateTicketDTO) error {
	ticket := &models.Ticket{
		ReporterID:         userID,
		ReporterRole:       role,
		ReportedSeekerID:   input.ReportedSeekerID,
		ReportedBusinessID: input.ReportedBusinessID,
		Type:               input.Type,
		Subject:            input.Subject,
		Description:        input.Description,
	}
	return s.Repo.CreateTicket(ticket)
}

func (s *TicketServices) GetMyTickets(userID uint) ([]models.Ticket, error) {
	return s.Repo.GetMyTickets(userID)
}

func (s *TicketServices) DeleteTicket(userID, ticketID uint) error {
	return s.Repo.DeleteTicket(ticketID, userID)
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