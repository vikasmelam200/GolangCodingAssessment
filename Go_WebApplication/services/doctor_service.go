package services

import (
	"Go_WebApplication/auth"
	"Go_WebApplication/config"
	"Go_WebApplication/logger"
	"Go_WebApplication/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all patients assigned to a doctor
func GetPatientsByDoctor(c *gin.Context) {
	// Extract doctor name from JWT token
	authInfo, err := auth.GetClaims(c)
	if err != nil {
		logger.Log.Error().Err(err).Msg("something went wrong")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "something went wrong"})
		return
	}

	doctorName := authInfo.Username
	//
	var patients []models.Patient
	if err := config.DB.Where("assigned_to = ?", doctorName).Find(&patients).Error; err != nil {
		logger.Log.Error().Err(err).Msg("No patients found")
		c.JSON(http.StatusNotFound, gin.H{"error": "No patients found"})
		return
	}
	c.JSON(http.StatusOK, patients)
}

// Update patient diagnosis
func UpdateDiagnosisByDoctor(c *gin.Context) {
	var patient models.Patient
	id := c.Param("id")

	if err := config.DB.First(&patient, id).Error; err != nil {
		logger.Log.Error().Err(err).Msg("No patients found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var updatedData struct {
		Diagnosis string `json:"diagnosis"`
	}

	if err := c.ShouldBindJSON(&updatedData); err != nil {
		logger.Log.Error().Err(err).Msg("While binding json : UpdateDiagonosisByDoctor")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient.Diagnosis = updatedData.Diagnosis
	config.DB.Save(&patient)
	c.JSON(http.StatusOK, gin.H{"message": "Diagnosis updated", "patient": patient})
}
