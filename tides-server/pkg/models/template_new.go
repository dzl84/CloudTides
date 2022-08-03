package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)


type TemplateNew struct {
	gorm.Model

	// templateID
	TemplateID uint `gorm:"uniqueIndex" json:"templateID,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// provider orgID
	ProvideOrgID uint `json:"provideOrgID,omitempty"`

	// VM template ID (VMTemp in previous template.go)
	VmTemplateID pq.Int64Array `json:"vmTemplateID,omitempty" gorm:"type:integer[]"`

	// type
	Type string `json:"type,omitempty"`


}