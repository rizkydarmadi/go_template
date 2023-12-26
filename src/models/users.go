package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name  string
	Email string
}
