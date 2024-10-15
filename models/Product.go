package models

import (
	"gorm.io/gorm"
)

type Product struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Name       string  `gorm:"not null;unique;size:50" json:"name"`
	Value      float64 `gorm:"not null;size:50" json:"value"`
	CategoryId uint    `json:"category_id"`
	gorm.Model
}
