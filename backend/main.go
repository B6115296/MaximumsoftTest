package main

import (
	"fmt"
	AuthController "jwt-api/controller/auth"
	EmployeeController "jwt-api/controller/employee"
	"jwt-api/middleware"
	"jwt-api/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Binding from JSON
type Register struct {
	Username string `json:"username" bingding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

// Binding from JSON
type RegisterEmployee struct {
	Idcard   string `json:"username" bingding:"required"`
	Fullname string `json:"password" binding:"required"`
	Phone    string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

type Admin struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Avatar   string
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	orm.InitDB()

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	authorized := r.Group("/employee", middleware.JWTAuthen())
	authorized.GET("/getemployees", EmployeeController.GetEmployees)
	authorized.GET("/getemployee", EmployeeController.GetEmployee)
	authorized.POST("/createemployee", EmployeeController.CreateEmployee)
	authorized.PUT("/updateemployee/:id", EmployeeController.UpdateEmployee)
	authorized.DELETE("/deleteemployee/:id", EmployeeController.DeleteEmployee)

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
