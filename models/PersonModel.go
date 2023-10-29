package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	ID        uint
	Address   string
	Phone     uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Person) TableName() string {
	return "persons"
}
