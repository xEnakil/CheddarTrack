package model

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey"`
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Transactions []Transaction
}

type Category struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Type string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Transactions []Transaction
}

type Transaction struct {
	ID uint `gorm:"primaryKey"`
	UserID uint `gorm:"not null"`
	CategoryID uint `gorm:"not null"`
	Amount float64 `gorm:"not null"`
	Currency string `gorm:"not null"`
	Description string
	Timestamp time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"not null"`

	User User
	Category Category
}

