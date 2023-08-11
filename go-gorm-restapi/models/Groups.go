package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model

	Name   string `gorm:"not null;unique_index" json:"name"`
	UserID uint   `json:"user_id"`
	Tasks  []Task `gorm:"many2many:group_tasks;"`
}
