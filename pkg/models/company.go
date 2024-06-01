package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string
	Email string
	Address string
	Abn string
	PhoneNumber string
	UserID uint
}