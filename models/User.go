package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string `gorm:"type:varchar(100);not null"`
	Lastname  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);not null; unique_index"`
	Task      []Task
}
