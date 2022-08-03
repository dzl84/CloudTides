package models

import (
	"gorm.io/gorm"
)


type ResVapp struct {
	gorm.Model

	// vapp_id
	VappID uint `json:"vappID,omitempty"`

	// template_id
	TemplateID uint `json:"templateID,omitempty"`

	// Res_id
	ResID uint `json:"resID,omitempty"`

}