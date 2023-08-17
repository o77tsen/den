package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Author  string   `json:"author" binding:"required"`
	Message string   `json:"message" binding:"required"`
	Answer  []string `json:"answer"`
}

type CreateMessageReq struct {
	gorm.Model
	Author  string   `json:"author" binding:"required"`
	Message string   `json:"message" binding:"required"`
	Answer  []string `json:"answer"`
}

type UpdateMessageReq struct {
	gorm.Model
	Author  string   `json:"author" binding:"required"`
	Message string   `json:"message" binding:"required"`
	Answer  []string `json:"answer"`
}
