package orm

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Idcard   string
	Fullname string
	Phone    string
	Avatar   string
}
