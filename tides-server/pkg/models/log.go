package models

import (
	"gorm.io/gorm"
	"time"
)


type Log struct {
	gorm.Model

	// user_id
	UserID uint `json:"userID,omitempty"`
	
	// operation
	Operation string `json:"operation,omitempty"`

	// Time
	Time time.Time `json:"time,omitempty"`

	// status
	Status string `json:"status,omitempty"`

}