package orm

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Avatar   string
}
