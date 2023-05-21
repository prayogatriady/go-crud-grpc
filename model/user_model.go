package model

import (
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	ID        int            `json:"user_id"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Balance   int            `json:"balance"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
