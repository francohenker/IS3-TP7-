package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"type:varchar(256);not null;unique_index"`
	Description string `gorm:"type:varchar(300)"`
	Done        bool   `gorm:"defaul:false; not null"`
	UserId      uint   `gorm:"not null"`
}
