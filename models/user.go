package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	EmailId   string    `json:"emailId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
