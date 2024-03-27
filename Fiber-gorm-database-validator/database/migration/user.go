package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint   `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreateAt  time.Time
	UpdateAt  time.Time
	DeletetAt gorm.DeletedAt `gorm:"index"`
}
