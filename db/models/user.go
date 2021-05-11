package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"username"`
	Name     string `json:"name"`
	Phone    string `json:"phone" gorm:"unique"`
	City     string `json:"city"`
}
