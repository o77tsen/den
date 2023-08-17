package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID         uint           `json:"id" gorm:"primary_key"`
	Author     string         `json:"author" binding:"required"`
	Message    string         `json:"message" binding:"required"`
	Answer     pq.StringArray `json:"answer" gorm:"type:varchar(1000)[]"`
	IsAnswered bool           `json:"is_answered"`
}

type CreateMessageReq struct {
	gorm.Model
	ID         uint           `json:"id" gorm:"primary_key"`
	Author     string         `json:"author" binding:"required"`
	Message    string         `json:"message" binding:"required"`
	Answer     pq.StringArray `json:"answer" gorm:"type:varchar(1000)[]"`
	IsAnswered bool           `json:"is_answered"`
}
