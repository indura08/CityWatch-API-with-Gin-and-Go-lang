package main

import (
	"citywatch/internal/database"
	"citywatch/internal/handler"
	"citywatch/internal/models"
	"citywatch/internal/repository"
	"citywatch/internal/routes"
	"citywatch/internal/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	database.ConnectToDB()
	err = database.DB.AutoMigrate(&models.Incident{}, &models.User{})

	if err != nil {
		log.Fatal("Migration failed")
	}

	// setup the initializers for Repo, service and handler

	userRepo := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	incidentRepo := repository.NewIncidentRepository(database.DB)
	incidentService := service.NewIncidentService(incidentRepo, userRepo)
	incidentHanlder := handler.NewIncidentHandler(&incidentService)

	//define other handlers,services and repos if available
	r := gin.Default()

	routes.AuthRoutes(r, userHandler)
	routes.IncidentRoutes(r, incidentHanlder)

	//starting the server
	//this is the simplest way
	// if err := r.Run(":8080"); err != nil {
	// 	log.Fatal("Server failed the error is: ", err.Error())
	// }

	//this is the more custom way to start the server
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
