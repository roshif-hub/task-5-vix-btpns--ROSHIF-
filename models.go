package main

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Photos    []Photo   `json:"photos" gorm:"foreignKey:UserID"`
}

type Photo struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (Photo) TableName() string {
	return "photos"
}

// ... your other model structures ...
