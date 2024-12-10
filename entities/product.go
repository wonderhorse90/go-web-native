package entities

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	CategoryID  uint      `gorm:"column:category_id" json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
