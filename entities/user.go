package entities

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:regular"` // Either "admin" or "regular"
}
