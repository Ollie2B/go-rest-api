package models

import "time"

type Song struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name   	  string    `json:"name"`
	Artist    string		`json:"artist"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateSong struct {
	Name  string  `json:"name" binding:"required"`
	Artist string `json:"artist" binding:"required"`
}

type UpdateSong struct {
	Name  string  `json:"name"`
	Artist string `json:"artist"`
}