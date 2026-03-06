
package controllers

import (
	"hoodhire/internal/services"
	"hoodhire/structures/dto"
	"hoodhire/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type JobController struct {
	Service *services.JobServices
}

func NewJobHandler(serv *services.JobServices) *JobController {
	return &JobController{Service: serv}
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Job CRUD~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (jc *JobController) CreateJob(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	input, err := utils.BindAndValidate[dto.CreateJobDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := jc.Service.CreateJob(userID, input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "job posted successfully"})
}

func (jc *JobController) GetJobByID(c fiber.Ctx) error {
	jobID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid job id"})
	}
	job, err := jc.Service.GetJobByID(uint(jobID))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "job not found"})
	}
	return c.Status(200).JSON(fiber.Map{"job": job})
}

func (jc *JobController) GetMyJobs(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	jobs, err := jc.Service.GetMyJobs(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"jobs": jobs})
}

func (jc *JobController) GetActiveJobs(c fiber.Ctx) error {
	jobs, err := jc.Service.GetActiveJobs()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"jobs": jobs})
}

func (jc *JobController) GetJobsByCategory(c fiber.Ctx) error {
	categoryID, err := strconv.ParseUint(c.Params("categoryID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid category id"})
	}
	jobs, err := jc.Service.GetJobsByCategory(uint(categoryID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"jobs": jobs})
}

func (jc *JobController) GetJobsByLocality(c fiber.Ctx) error {
	locality := c.Params("locality")
	if locality == "" {
		return c.Status(400).JSON(fiber.Map{"error": "locality is required"})
	}
	jobs, err := jc.Service.GetJobsByLocality(locality)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"jobs": jobs})
}

func (jc *JobController) UpdateJob(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	jobID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid job id"})
	}
	input, err := utils.BindAndValidate[dto.UpdateJobDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := jc.Service.UpdateJob(userID, uint(jobID), input); err != nil {
		if err.Error() == "unauthorized" {
			return c.Status(403).JSON(fiber.Map{"error": "you can only edit your own jobs"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "job updated successfully"})
}

func (jc *JobController) UpdateJobStatus(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	jobID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid job id"})
	}
	input, err := utils.BindAndValidate[dto.UpdateJobStatusDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := jc.Service.UpdateJobStatus(userID, uint(jobID), input); err != nil {
		if err.Error() == "unauthorized" {
			return c.Status(403).JSON(fiber.Map{"error": "you can only update your own jobs"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "job status updated successfully"})
}

func (jc *JobController) DeleteJob(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	jobID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid job id"})
	}
	if err := jc.Service.DeleteJob(userID, uint(jobID)); err != nil {
		if err.Error() == "unauthorized" {
			return c.Status(403).JSON(fiber.Map{"error": "you can only delete your own jobs"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "job deleted successfully"})
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Applications~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (jc *JobController) ApplyToJob(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	jobID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid job id"})
	}
	input, err := utils.BindAndValidate[dto.JobApplicationDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := jc.Service.ApplyToJob(userID, uint(jobID), input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "application submitted successfully"})
}

func (jc *JobController) GetApplicationsForJob(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	jobID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid job id"})
	}
	applications, err := jc.Service.GetApplicationsForJob(userID, uint(jobID))
	if err != nil {
		if err.Error() == "unauthorized" {
			return c.Status(403).JSON(fiber.Map{"error": "you can only view applications for your own jobs"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"applications": applications})
}

func (jc *JobController) GetMyApplications(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	applications, err := jc.Service.GetMyApplications(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"applications": applications})
}

func (jc *JobController) UpdateApplicationStatus(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	applicationID, err := strconv.ParseUint(c.Params("applicationID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid application id"})
	}
	input, err := utils.BindAndValidate[dto.UpdateApplicationStatusDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := jc.Service.UpdateApplicationStatus(userID, uint(applicationID), input); err != nil {
		if err.Error() == "unauthorized" {
			return c.Status(403).JSON(fiber.Map{"error": "unauthorized"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "application status updated successfully"})
}

func (jc *JobController) WithdrawApplication(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	applicationID, err := strconv.ParseUint(c.Params("applicationID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid application id"})
	}
	if err := jc.Service.WithdrawApplication(userID, uint(applicationID)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "application withdrawn successfully"})
}