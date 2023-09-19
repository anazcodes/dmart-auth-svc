package models

import (
	"time"
)

type Model struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type User struct {
	Model
	Username string `gorm:"not null"`
	Email    string //can login with email or phone, nullable field
	Phone    int64  //can login with email or phone, nullable fields
	Password string `gorm:"not null"`
}

type Address struct {
	Model
	Name             string `gorm:"not null"`
	PhoneNumber      string `gorm:"not null"`
	PostalCode       string `gorm:"not null"`
	AddressLine      string `gorm:"not null"`
	Locality         string
	District         string
	Landmark         string
	AlternativePhone string
	UserID           uint `gorm:"not null"`
	User             User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IsDefault        bool `gorm:"default:false"`
}
