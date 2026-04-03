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

func (s *TicketServices)ResolveTicket(id uint,reply string)error{
	return s.Repo.UpdateTicketStatus(id,"resolved",reply)
}
func (s *TicketServices) ReviewTicket(id uint, reply string) error {
	return s.Repo.UpdateTicketStatus(id, "reviewed", reply)
}

func (s *TicketServices) DismissTicket(id uint,reply string) error {
	var response string
	if reply!=""{
		 response = reply}else{
			response=""
		 }
	return s.Repo.UpdateTicketStatus(id, "dismissed",response)
}
func ( s *TicketServices)GetTickets(status,tType string, page,limit int)([]models.Ticket,error){
	offset := (page - 1) * limit
	return s.Repo.GetTickets(status,tType,limit,offset)
}

func (s *TicketServices) GetTicketsByBusiness(businessID uint) ([]models.Ticket, error) {
	return s.Repo.GetTicketsByBusiness(businessID)
}