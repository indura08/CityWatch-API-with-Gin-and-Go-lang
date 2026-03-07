package handler

import (
	"citywatch/internal/DTO/Incident"
	"citywatch/internal/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IncidentHandler struct {
	incidentService service.IncidentService
}

func NewIncidentHandler(i *service.IncidentService) *IncidentHandler {
	return &IncidentHandler{incidentService: i}
}

func (h *IncidentHandler) CreateIncident(c *gin.Context) {
	var incidentDto *Incident.IncidentDto

	if err := c.ShouldBindJSON(&incidentDto); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	var imagePath string
	if incidentDto.Image != nil{
		fileName := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), incidentDto.Image.Filename)

		err2 := c.SaveUploadedFile(incidentDto.Image, fileName)
		if err2 != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Image uplaod fails, check again",
			})
			return
		}

		imagePath = fileName
	}

	err3 := h.incidentService.CreateIncident(incidentDto, imagePath)
	if err3 != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err3.Error()
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Incident Created Successfully"
	})
}

//methn indla hnna thiynne 