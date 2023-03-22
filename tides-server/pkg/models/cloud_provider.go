package models

import (
	"gorm.io/gorm"
)


type CloudProvider struct {
	gorm.Model

	// cloud_provider_id
	CloudProviderID uint `gorm:"uniqueIndex" json:"cloudPrividerID,omitempty"`

	// org_id
	OrgID uint `json:"orgID,omitempty" sql:"type:uint REFERENCES Org(ID)"`

	// org
	Org Org `gorm:"ForeignKey:OrgID;AssociationForeignKey:ID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// cp_backing_type
	CpBackingType string `json:"cpBackingType,omitempty"`

}