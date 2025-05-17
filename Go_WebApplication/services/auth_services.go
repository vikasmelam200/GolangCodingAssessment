package services

import (
	"Go_WebApplication/auth"
	"Go_WebApplication/config"
	"Go_WebApplication/logger"
	"Go_WebApplication/models"
	"Go_WebApplication/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Signup Handler with Validation
func Signup(c *gin.Context) {
	var request models.User
	// Bind and validate JSON input
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Log.Error().Err(err).Msg("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate role (only "receptionist" or "doctor" allowed)
	validRoles := map[string]bool{"receptionist": true, "doctor": true}

	if _, exists := validRoles[strings.ToLower(request.RoleName)]; !exists {
		logger.Log.Error().Msg("Invalid role. Allowed roles: receptionist, doctor")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role. Allowed roles: receptionist, doctor"})
		return
	}

	// Check if the username already exists
	var existingUser models.User
	if err := config.DB.Where("username = ?", request.Username).First(&existingUser).Error; err == nil {
		logger.Log.Error().Err(err).Msg("Username already exists try new")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists try new"})
		return
	}

	// Validate password
	if err := util.ValidatePassword(request.Password); err != nil {
		logger.Log.Error().Err(err).Msg("Invalid password")
		c.JSON(http.StatusBadRequest, gin.H{"Invalid password": err.Error})
		return
	}

	// Hash the password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	request.Password = string(hashedPassword)

	// Save user to the database
	err := config.DB.Create(&request).Error
	if err != nil {
		logger.Log.Error().Err(err).Msg("error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully", 
	})
}

// Login Handler
func Login(c *gin.Context) {
	//
	var req models.LoginUser
	var response models.User
	var loginResponse models.LoginResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error().Err(err).Msg("error while binding json: Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user
	if err := config.DB.Where("username = ?", req.Username).First(&response).Error; err != nil {
		logger.Log.Error().Err(err).Msg("invalid username")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(req.Password)); err != nil {
		logger.Log.Error().Err(err).Msg("Invalid password")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// token generate
	token, err := auth.GenerateJWT(response.Username, response.Email, response.RoleName)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//
	loginResponse.Username = response.Username
	loginResponse.Email = response.Email
	loginResponse.RoleName = response.RoleName
	loginResponse.Token = token

	c.JSON(http.StatusOK, gin.H{"login succussfully": loginResponse})
}
