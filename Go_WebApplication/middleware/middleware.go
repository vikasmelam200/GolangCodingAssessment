package middleware

import (
	"Go_WebApplication/auth"
	"Go_WebApplication/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Restrict access to Receptionists
func ReceptionistOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		authInfo, err := auth.GetClaims(c)
		if err != nil {
			logger.Log.Error().Err(err).Msg("something went wrong")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "something went wrong"})
			return
		}
		if authInfo.RoleName != "receptionist" {
			logger.Log.Error().Err(err).Msg("Access restricted to receptionists")
			c.JSON(http.StatusForbidden, gin.H{"error": "Access restricted to receptionists"})
			c.Abort()
		}
	}
}

// Restrict access to Doctors
func DoctorOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		authInfo, err := auth.GetClaims(c)
		if err != nil {
			logger.Log.Error().Err(err).Msg("something went wrong")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "something went wrong"})
			return
		}
		if authInfo.RoleName != "doctor" {
			logger.Log.Error().Err(err).Msg("Access restricted to doctor")
			c.JSON(http.StatusForbidden, gin.H{"error": "Access restricted to doctor"})
			c.Abort()
		}
	}
}
