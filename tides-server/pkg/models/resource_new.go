package models

import (
	"gorm.io/gorm"
)


type ResourceNew struct {
	gorm.Model

	// res_id
	ResID uint `gorm:"uniqueIndex" json:"resID,omitempty"`

	// org_id
	OrgID uint `json:"orgID,omitempty"`

	// res_path
	ResPath string `json:"resPath,omitempty"`

	// res_status
	ResStatus string `json:"resStatus,omitempty"`

	// cloudprovider_id
	CloudProviderID uint `json:"cloudProviderID,omitempty" sql:"type:uint REFERENCES CloudProvider(CloudProviderID)"`

	// cloudProvider
	CloudProvider CloudProvider `gorm:"ForeignKey:CloudProviderID;AssociationForeignKey:CloudProviderID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

}