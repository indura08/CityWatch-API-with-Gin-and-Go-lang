package handler

import (
	"citywatch/internal/DTO/Incident"
	"citywatch/internal/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type IncidentHandler struct {
	incidentService *service.IncidentService
}

func NewIncidentHandler(i *service.IncidentService) *IncidentHandler {
	return &IncidentHandler{incidentService: i}
}

func (h *IncidentHandler) CreateIncident(c *gin.Context) {
	var incidentDto *Incident.IncidentDto

	if err := c.ShouldBindJSON(&incidentDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	var imagePath string
	if incidentDto.Image != nil {
		fileName := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), incidentDto.Image.Filename)

		err2 := c.SaveUploadedFile(incidentDto.Image, fileName)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Image uplaod fails, check again",
			})
			return
		}

		imagePath = fileName
	}

	err3 := h.incidentService.CreateIncident(incidentDto, imagePath)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err3.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Incident Created Successfully",
	})
}

func (h *IncidentHandler) GetAllIncidents(c *gin.Context) {
	incidentList, err1 := h.incidentService.GetAllIncidents()
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err1.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully fetched incidents",
		"data":    incidentList,
	})
}

func (h *IncidentHandler) DeleteIncidentByID(c *gin.Context) {
	incidentId, err1 := strconv.Atoi(c.Param("incidentId"))
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid incident ID",
		})
		return
	}

	err2 := h.incidentService.DeleteIncidentById(incidentId)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err2.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "incident deleted successfully",
	})
}

func (h *IncidentHandler) AssignWorkerToIncident(c *gin.Context) {
	var assignWorkerDto *Incident.AssignWorkerIncidentDto

	if err1 := c.ShouldBindJSON(&assignWorkerDto); err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	err2 := h.incidentService.AssignWorkerToIncident(assignWorkerDto.IncidentId, assignWorkerDto.WorkerId)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err2.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Worker assigned successfully",
	})
}

func (h *IncidentHandler) ChangeIncidentStatus(c *gin.Context) {
	var incidentStatusChangeDto *Incident.IncidentStatusChangeDto
	if err1 := c.ShouldBindJSON(&incidentStatusChangeDto); err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	incidentId, err := strconv.Atoi(c.Param("incidentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Incident ID",
		})
		return
	}

	err2 := h.incidentService.UpdateIncidentStatus(incidentId, incidentStatusChangeDto.IncidentStatus)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err2.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Incident status updated succesfully",
	})
}

func (h *IncidentHandler) UpdateIncident(c *gin.Context) {
	var incidentDto *Incident.IncidentDto
	if err1 := c.ShouldBindJSON(&incidentDto); err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request Body",
		})
		return
	}

	incidentId, err := strconv.Atoi(c.Param("incidentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid incident id",
		})
		return
	}

	err2 := h.incidentService.UpdateIncident(incidentId, incidentDto)
	if err2 != nil {
		c.JSON(500, gin.H{
			"error": err2.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Incident updated succesfully",
	})
}
