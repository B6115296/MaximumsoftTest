package employee

import (
	"jwt-api/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Binding from JSON
type EmployeeBody struct {
	Idcard   string
	Fullname string
	Phone    string
	Avatar   string
}

func CreateEmployee(c *gin.Context) {
	var json EmployeeBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check employee Exists
	var idcardExists orm.Employee
	orm.Db.Where("Idcard = ?", json.Idcard).First(&idcardExists)
	if idcardExists.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Idcard Exists"})
		return
	}

	// Create employee
	employee := orm.Employee{Idcard: json.Idcard, Fullname: json.Fullname, Phone: json.Phone, Avatar: json.Avatar}
	orm.Db.Create(&employee)
	if employee.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":     "ok",
			"message":    "Employee Create Success",
			"employeeId": employee.ID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "employee Create Failed", "employeeId": employee.ID})
	}
}

func GetEmployees(c *gin.Context) {
	var employees []orm.Employee
	orm.Db.Find(&employees)
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"message":  "Employee Detail",
		"employee": employees,
	})
}

func GetEmployee(c *gin.Context) {
	var employees []orm.Employee
	orm.Db.Find(&employees)
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"message":  "Employee Detail",
		"employee": employees,
	})
}

func UpdateEmployee(c *gin.Context) {
	var json EmployeeBody
	idcard := c.Param("id")
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var updatedEmployee orm.Employee
	orm.Db.First(&idcard)
	updatedEmployee.Idcard = idcard
	updatedEmployee.Fullname = json.Fullname
	updatedEmployee.Phone = json.Phone
	updatedEmployee.Avatar = json.Avatar
	orm.Db.Save(&updatedEmployee)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Employee Update Success", "employeeId": updatedEmployee.Idcard})
}

func DeleteEmployee(c *gin.Context) {
	idcard := c.Param("id")
	var employee orm.Employee
	orm.Db.Where("idcard = ?", idcard).Delete(&employee)
	if employee.Idcard == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Employee Delete Success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "employee Delete Failed"})
	}
}
