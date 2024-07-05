package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string     `json:"name" gorm:"required"`
	Email    string     `json:"email" gorm:"index:email_1,unique,required"`
	Password string     `json:"password" gorm:"required"`
	DOB      *time.Time `json:"date"`
}
