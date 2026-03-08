package routes

import (
	"citywatch/internal/handler"
	"citywatch/internal/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func IncidentRoutes(r *gin.Engine, incidentHanlder *handler.IncidentHandler) {
	incidentRoute := r.Group("/api/incident/")

	incidentRoute.GET("/incidents", middleware.AuthorizeRoles(os.Getenv("SECRET"), 0, 1, 2), incidentHanlder.GetAllIncidents)
	incidentRoute.POST("/create", middleware.AuthorizeRoles(os.Getenv("SECRET"), 0, 1, 2), incidentHanlder.CreateIncident)
	incidentRoute.DELETE("/delete/:incidentId", middleware.AuthorizeRoles(os.Getenv("SECRET"), 0, 1, 2), incidentHanlder.DeleteIncidentByID)
	incidentRoute.PUT("/assignWorker", middleware.AuthorizeRoles(os.Getenv("SECRET"), 0, 2), incidentHanlder.AssignWorkerToIncident)
	incidentRoute.PUT("/changeStatus", middleware.AuthorizeRoles(os.Getenv("SECRET"), 0, 2), incidentHanlder.ChangeIncidentStatus)
	incidentRoute.PUT("/update", middleware.AuthorizeRoles(os.Getenv("SECRET"), 0, 1, 2), incidentHanlder.UpdateIncident)
}
