package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID      string `gorm:"index;unique;type:varchar(22)" json:"user_id"`
	Name        string `gorm:"type:varchar(30)" json:"name"`
	Email       string `gorm:"unique;email" json:"email"`
	Password    string `json:"-"`
	PhoneNumber string `gorm:"unique;type:varchar(30)" json:"phone_number"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `gorm:"-" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
