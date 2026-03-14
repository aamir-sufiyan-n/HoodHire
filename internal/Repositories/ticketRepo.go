package repositories

import (
	"hoodhire/structures/models"

	"gorm.io/gorm"
)

type TicketRepo struct {
    DB *gorm.DB
}

func (r *TicketRepo) CreateTicket(ticket *models.Ticket) error {
    ticket.Status = "open"
    return r.DB.Create(ticket).Error
}

func (r *TicketRepo) GetMyTickets(seekerID uint) ([]models.Ticket, error) {
    var tickets []models.Ticket
    err := r.DB.Preload("Business").
        Where("seeker_id = ?", seekerID).
        Find(&tickets).Error
    return tickets, err
}

func (r *TicketRepo) DeleteTicket(ticketID, seekerID uint) error {
    return r.DB.Unscoped().
        Where("id = ? AND seeker_id = ?", ticketID, seekerID).
        Delete(&models.Ticket{}).Error
}

// admin
func (r *TicketRepo) GetAllTickets() ([]models.Ticket, error) {
    var tickets []models.Ticket
    err := r.DB.Preload("Business").Preload("Seeker").
        Find(&tickets).Error
    return tickets, err
}

func (r *TicketRepo) GetTicketsByType(ticketType string) ([]models.Ticket, error) {
    var tickets []models.Ticket
    err := r.DB.Preload("Business").Preload("Seeker").
        Where("type = ?", ticketType).
        Find(&tickets).Error
    return tickets, err
}

func (r *TicketRepo) GetTicketsByStatus(status string) ([]models.Ticket, error) {
    var tickets []models.Ticket
    err := r.DB.Preload("Business").Preload("Seeker").
        Where("status = ?", status).
        Find(&tickets).Error
    return tickets, err
}

func (r *TicketRepo) UpdateTicketStatus(ticketID uint, status string) error {
    return r.DB.Model(&models.Ticket{}).
        Where("id = ?", ticketID).
        Update("status", status).Error
}

func (r *TicketRepo) GetTicketsByBusiness(businessID uint) ([]models.Ticket, error) {
    var tickets []models.Ticket
    err := r.DB.Preload("Seeker").
        Where("business_id = ?", businessID).
        Find(&tickets).Error
    return tickets, err
}