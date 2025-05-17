package routes

import (
	"Go_WebApplication/auth"
	"Go_WebApplication/middleware"
	"Go_WebApplication/services"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes all routes
func SetupRoutes() *gin.Engine {
	r := gin.Default()
	// Authentication Routes
	gr := r.Group("/api/auth/v1")
	gr.POST("/signup", services.Signup)
	gr.POST("/login", services.Login)

	// Protected Receptionist Routes
	receptionist := r.Group("/api/receptionist/v1")
	receptionist.Use(auth.Auth())
	receptionist.Use(middleware.ReceptionistOnly())
	{
		receptionist.POST("/add-patient", services.CreatePatient)
		receptionist.GET("/patients", services.GetAllPatients)
		receptionist.GET("/patient/:id", services.GetPatientByID)
		receptionist.PUT("/patient/:id", services.UpdatePatient)
		receptionist.DELETE("/patient/:id", services.DeletePatient)
	}

	// Doctor Routes (Restricted Access)
	doctor := r.Group("api/doctor/v1")
	doctor.Use(auth.Auth())
	doctor.Use(middleware.DoctorOnly())
	{
		doctor.GET("/patients", services.GetPatientsByDoctor)        // View assigned patients
		doctor.PUT("/patient/:id", services.UpdateDiagnosisByDoctor) // Update diagnosis
	}

	return r
}
