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

func (r *TicketRepo) GetMyTickets(userID uint) ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := r.DB.Preload("ReportedBusiness").Preload("ReportedSeeker").
		Where("reporter_id = ?", userID).
		Find(&tickets).Error
	return tickets, err
}

func (r *TicketRepo) DeleteTicket(ticketID, userID uint) error {
	return r.DB.Unscoped().
		Where("id = ? AND reporter_id = ?", ticketID, userID).
		Delete(&models.Ticket{}).Error
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~admin~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func (r *TicketRepo) GetTickets(status, tType string, limit, offset int) ([]models.Ticket, error) {
	var tickets []models.Ticket

	query := r.DB.Preload("Reporter")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if tType != "" {
		query = query.Where("type = ?", tType)
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	err := query.Find(&tickets).Error
	return tickets, err
}

func (r *TicketRepo) UpdateTicketStatus(ticketID uint, status,reply string) error {
    	updates := map[string]interface{}{
		"status": status,
	}
	if reply != "" {
		updates["reply"] = reply
	}
	return r.DB.Model(&models.Ticket{}).
		Where("id = ?", ticketID).
		Updates(updates).Error
}

func (r *TicketRepo) GetTicketsByBusiness(businessID uint) ([]models.Ticket, error) {
    var tickets []models.Ticket
    err := r.DB.Preload("Seeker").
        Where("business_id = ?", businessID).
        Find(&tickets).Error
    return tickets, err
}