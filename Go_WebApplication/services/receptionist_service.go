package services

import (
	"Go_WebApplication/config"
	"Go_WebApplication/logger"
	"Go_WebApplication/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a new patient
func CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		logger.Log.Error().Err(err).Msg("While binding json : CreatePatient")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&patient)
	c.JSON(http.StatusCreated, gin.H{"message": "Patient added", "patient": patient})
}

// Get all patients
func GetAllPatients(c *gin.Context) {
	var patients []models.Patient
	err := config.DB.Find(&patients).Error
	if err != nil {
		logger.Log.Error().Err(err).Msg("error while getting records:")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, patients)
}

// Get a single patient by ID
func GetPatientByID(c *gin.Context) {
	var patient models.Patient
	id := c.Param("id")

	if err := config.DB.First(&patient, id).Error; err != nil {
		logger.Log.Error().Err(err).Msg("while getting patient data")
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

// Update patient details
func UpdatePatient(c *gin.Context) {
	var patient models.Patient
	id := c.Param("id")

	if err := config.DB.First(&patient, id).Error; err != nil {
		logger.Log.Error().Err(err).Msg("while getting patient data")
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	if err := c.ShouldBindJSON(&patient); err != nil {
		logger.Log.Error().Err(err).Msg("while binding json data :")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := config.DB.Save(&patient).Error
	if err != nil {
		logger.Log.Error().Err(err).Msg("while updating patient data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient updated", "patient": patient})
}

// Delete a patient
func DeletePatient(c *gin.Context) {
	var patient models.Patient
	id := c.Param("id")

	if err := config.DB.First(&patient, id).Error; err != nil {
		logger.Log.Error().Err(err).Msg("while getting patient data")
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	err := config.DB.Delete(&patient).Error
	if err != nil {
		logger.Log.Error().Err(err).Msg("while deleting patient data")
		c.JSON(http.StatusNotFound, gin.H{"error": "while deleting patient data"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}
