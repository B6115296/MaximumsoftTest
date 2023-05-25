package auth

import (
	"fmt"
	"jwt-api/orm"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret []byte

// Binding from JSON
type RegisterBody struct {
	Username string `json:"username" bingding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check Admin Exists
	var adminExist orm.Admin
	orm.Db.Where("username = ?", json.Username).First(&adminExist)
	if adminExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Admin Exists"})
		return
	}

	// Create Admin
	encrptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	admin := orm.Admin{Username: json.Username, Password: string(encrptedPassword),
		Fullname: json.Fullname, Avatar: json.Avatar}
	orm.Db.Create(&admin)
	if admin.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Admin Create Success",
			"adminId": admin.ID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Admin Create Failed", "adminId": admin.ID})
	}
}

// Binding from JSON
type LoginBody struct {
	Username string `json:"username" bingding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check Admin Exists
	var adminExist orm.Admin
	orm.Db.Where("username = ?", json.Username).First(&adminExist)
	if adminExist.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Admin Does not Exists"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(adminExist.Password), []byte(json.Password))
	if err == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"adminId": "adminExist.ID",
			"exp":     time.Now().Add(time.Minute * 5).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)

		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Login Success", "token": tokenString})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Login Failed"})
	}
}
