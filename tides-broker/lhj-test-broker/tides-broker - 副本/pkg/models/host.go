package models

import (
	"gorm.io/gorm"
)


type Host struct {
	gorm.Model

	// hostname
	Hostname string `gorm:"unique;not null"`

	// policy
    Policy int64 `json:"policy,omitempty"`

	// datacenter
    Datacenter string `json:"datacenters,omitempty"`

	// cluster
	Cluster string `json:"cluster,omitempty"`
    
	// ip
	IP string `json:"operation,omitempty"`

	// port
    Port int64 `json:"port,omitempty"`
	
	// sshkey
	Sshkey string `json:"sshkey,omitempty"`
           
	// username
    Username string `json:"username,omitempty"`

	// password
	Password string `json:"password,omitempty"`
}