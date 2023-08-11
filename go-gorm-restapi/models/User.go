package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string  `jgorm:"not null" json:"username"`
	Password string  `gorm:"not null" json:"password"`
	FirsName string  `gorm:"not null" json:"first_name"`
	LastName string  `gorm:"not null" json:"last_name"`
	Email    string  `gorm:"not null;uniqueIndex" json:"email"`
	Groups   []Group `json:"groups"`
}
