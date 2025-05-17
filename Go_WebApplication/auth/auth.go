package auth

import (
	"Go_WebApplication/config"
	"errors"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AuthInfo :
type AuthInfo struct {
	RoleName string `json:"rolename"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var jwtKey = []byte(config.GetAPISecret())

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleName string `json:"rolename"`
	jwt.StandardClaims
}

// GenerateJWT :
func GenerateJWT(userName string, email string, roleName string) (tokenString string, err error) {
	// expirationTime := time.Now().Add(24 * time.Hour)
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &JWTClaim{
		Email:    email,
		RoleName: roleName,
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

// Auth :
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := ValidateToken(context, tokenString)
		if err != nil {
			context.JSON(401, gin.H{"while validating token": err.Error})
			context.Abort()
			return
		}
		context.Next()
	}
}

// ValidateToken :
func ValidateToken(context *gin.Context, signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	//
	if claims.Username == "" || claims.Email == "" || claims.RoleName == "" {
		err = errors.New("token not valid")
		return
	}
	context.Set("claims", claims)
	return
}

func GetClaims(c *gin.Context) (AuthInfo, error) {
	var objAuthInfo AuthInfo

	claims, _ := c.Get("claims")
	//
	if claimsMap, ok := claims.(*JWTClaim); ok {
		objAuthInfo.RoleName = claimsMap.RoleName
		objAuthInfo.Username = claimsMap.Username
		objAuthInfo.Email = claimsMap.Email
	}

	return objAuthInfo, nil
}
