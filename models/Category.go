package models

type Category struct {
	ID      uint      `gorm:"primaryKey" json:"id"`
	Name    string    `gorm:"size:50; unique" json:"name"`
	Product []Product `gorm:"foreignKey:CategoryId"`
}
