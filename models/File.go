package models

type File struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Extension string `gorm:"not null"`
	Size      int64  `gorm:"not null"`
	Content   []byte `gorm:"not null"`
}
