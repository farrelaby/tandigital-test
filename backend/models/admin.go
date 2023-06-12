package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint   `json:"id" gorm:"primary_key;not null;auto_increment"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
