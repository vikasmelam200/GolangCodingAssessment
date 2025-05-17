package models

import (
	"gorm.io/gorm"
)

// data provided while sing-up :
type User struct {
	gorm.Model
	Username string `json:"username" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" gorm:"unique"`
	PhoneNo  string `json:"phone_no" gorm:"unique"`
	RoleName string `json:"role" binding:"required"` // "receptionist" or "doctor"
}

// login credentials :
type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse :
type LoginResponse struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleName string `json:"role_name"`
	Token    string `json:"token"`
}

// patient :
type Patient struct {
	gorm.Model
	Name       string `json:"name" binding:"required"`
	Age        int    `json:"age" binding:"required"`
	Gender     string `json:"gender" binding:"required"`
	Contact    string `json:"contact" gorm:"unique"`
	Diagnosis  string `json:"diagnosis"`
	AssignedTo string `json:"assigned_to"` // Doctor's Name
}
