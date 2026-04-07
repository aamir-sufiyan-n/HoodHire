package utils

import (
	"fmt"
	"hoodhire/structures/models"

	"github.com/jung-kurt/gofpdf"
)

func GenerateUsersPDF(users []models.User) (*gofpdf.Fpdf, error) {
    pdf := gofpdf.New("L", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(0, 10, "HoodHire - Users Report")
    pdf.Ln(12)

    // table header
    pdf.SetFont("Arial", "B", 11)
    pdf.SetFillColor(34, 197, 94)
    pdf.SetTextColor(255, 255, 255)
    pdf.CellFormat(10, 8, "ID", "1", 0, "C", true, 0, "")
    pdf.CellFormat(50, 8, "Username", "1", 0, "C", true, 0, "")
    pdf.CellFormat(80, 8, "Email", "1", 0, "C", true, 0, "")
    pdf.CellFormat(40, 8, "Role", "1", 0, "C", true, 0, "")
    pdf.CellFormat(30, 8, "Blocked", "1", 0, "C", true, 0, "")
    pdf.Ln(-1)

    // table rows
    pdf.SetFont("Arial", "", 10)
    pdf.SetTextColor(0, 0, 0)
    for i, u := range users {
        if i%2 == 0 {
            pdf.SetFillColor(240, 240, 240)
        } else {
            pdf.SetFillColor(255, 255, 255)
        }
        blocked := "No"
        if u.IsBlocked {
            blocked = "Yes"
        }
        pdf.CellFormat(10, 7, fmt.Sprintf("%d", u.ID), "1", 0, "C", true, 0, "")
        pdf.CellFormat(50, 7, u.Username, "1", 0, "L", true, 0, "")
        pdf.CellFormat(80, 7, u.Email, "1", 0, "L", true, 0, "")
        pdf.CellFormat(40, 7, u.Role, "1", 0, "C", true, 0, "")
        pdf.CellFormat(30, 7, blocked, "1", 0, "C", true, 0, "")
        pdf.Ln(-1)
    }
    return pdf, nil
}

func GenerateJobsPDF(jobs []models.Job) (*gofpdf.Fpdf, error) {
    pdf := gofpdf.New("L", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(0, 10, "HoodHire - Jobs Report")
    pdf.Ln(12)

    pdf.SetFont("Arial", "B", 11)
    pdf.SetFillColor(34, 197, 94)
    pdf.SetTextColor(255, 255, 255)
    pdf.CellFormat(10, 8, "ID", "1", 0, "C", true, 0, "")
    pdf.CellFormat(80, 8, "Title", "1", 0, "C", true, 0, "")
    pdf.CellFormat(60, 8, "Business", "1", 0, "C", true, 0, "")
    pdf.CellFormat(30, 8, "Status", "1", 0, "C", true, 0, "")
    pdf.CellFormat(30, 8, "Type", "1", 0, "C", true, 0, "")
    pdf.Ln(-1)

    pdf.SetFont("Arial", "", 10)
    pdf.SetTextColor(0, 0, 0)
    for i, j := range jobs {
        if i%2 == 0 {
            pdf.SetFillColor(240, 240, 240)
        } else {
            pdf.SetFillColor(255, 255, 255)
        }
        title := ""
        jobType := ""
        if j.Description != nil {
            title = j.Description.Title
            jobType = j.Description.JobType
        }
        businessName := ""
        if j.Business.ID != 0 {
            businessName = j.Business.BusinessName
        }
        pdf.CellFormat(10, 7, fmt.Sprintf("%d", j.ID), "1", 0, "C", true, 0, "")
        pdf.CellFormat(80, 7, title, "1", 0, "L", true, 0, "")
        pdf.CellFormat(60, 7, businessName, "1", 0, "L", true, 0, "")
        pdf.CellFormat(30, 7, j.Status, "1", 0, "C", true, 0, "")
        pdf.CellFormat(30, 7, jobType, "1", 0, "C", true, 0, "")
        pdf.Ln(-1)
    }
    return pdf, nil
}

func GenerateBusinessesPDF(businesses []models.Business) (*gofpdf.Fpdf, error) {
    pdf := gofpdf.New("L", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(0, 10, "HoodHire - Businesses Report")
    pdf.Ln(12)

    pdf.SetFont("Arial", "B", 11)
    pdf.SetFillColor(34, 197, 94)
    pdf.SetTextColor(255, 255, 255)
    pdf.CellFormat(10, 8, "ID", "1", 0, "C", true, 0, "")
    pdf.CellFormat(80, 8, "Business Name", "1", 0, "C", true, 0, "")
    pdf.CellFormat(50, 8, "Locality", "1", 0, "C", true, 0, "")
    pdf.CellFormat(30, 8, "Verified", "1", 0, "C", true, 0, "")
    pdf.Ln(-1)

    pdf.SetFont("Arial", "", 10)
    pdf.SetTextColor(0, 0, 0)
    for i, b := range businesses {
        if i%2 == 0 {
            pdf.SetFillColor(240, 240, 240)
        } else {
            pdf.SetFillColor(255, 255, 255)
        }
        verified := "No"
        if b.IsVerified {
            verified = "Yes"
        }
        pdf.CellFormat(10, 7, fmt.Sprintf("%d", b.ID), "1", 0, "C", true, 0, "")
        pdf.CellFormat(80, 7, b.BusinessName, "1", 0, "L", true, 0, "")
        pdf.CellFormat(50, 7, b.Locality, "1", 0, "C", true, 0, "")
        pdf.CellFormat(30, 7, verified, "1", 0, "C", true, 0, "")
        pdf.Ln(-1)
    }
    return pdf, nil
}