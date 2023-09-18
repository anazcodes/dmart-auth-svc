package models

import (
	"time"
)

type model struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type User struct {
	model
	username string `gorm:"not null"`
	email    string //can login with email or phone, nullable field
	phone    int64  //can login with email or phone, nullable fields
	password string `gorm:"not null"`
}

type Address struct {
	model
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
