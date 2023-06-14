package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `json:"id" gorm:"primary_key;not null;auto_increment"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Password  string         `json:"password" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Voucher struct {
	ID         uint      `json:"id" gorm:"primary_key;not null;auto_increment"`
	Name       string    `json:"name" gorm:"not null"`
	Code       string    `json:"code" gorm:"unique;not null"`
	Value      int       `json:"value" gorm:"not null"`
	Price      int       `json:"price" gorm:"not null"`
	ExpiryDate time.Time `json:"expiry_date" gorm:"not null"`
	Status     string    `json:"status" gorm:"not null"`
	// Category     string         `json:"category" gorm:"not null"`
	Quantity     int            `json:"quantity" gorm:"not null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Transactions []Transaction  `gorm:"many2many:transaction_voucher;"`
}

type Consumer struct {
	ID           uint           `json:"id" gorm:"primary_key;not null;auto_increment"`
	Email        string         `json:"email" gorm:"unique;not null"`
	Password     string         `json:"password" gorm:"not null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Transactions []Transaction  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Transaction struct {
	ID         uint           `json:"id" gorm:"primary_key;not null;auto_increment"`
	ConsumerID uint           `json:"consumer_id" gorm:"not null"`
	Status     string         `json:"status" gorm:"not null"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Vouchers   []Voucher      `gorm:"many2many:transaction_voucher;"`
}

// type TransactionVoucher struct {
// 	TransactionID uint `json:"transaction_id" gorm:"not null"`
// 	VoucherID     uint `json:"voucher_id" gorm:"not null"`
// 	// Quantity      int  `json:"quantity" gorm:"not null"`
// }
