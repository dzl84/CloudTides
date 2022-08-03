package models

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	Link          string `json:"link,omitempty"`
	ContainerID   string `json:"containerID,omitempty"`
	ContainerName string `json:"containerName,omitempty"`
	InstanceName  string `json:"instanceName,omitempty"`
	Token         string `json:"token,omitempty" gorm:"unique"`
	Uid           string `json:"uid,omitempty" gorm:"index"`
	SshHost       string `json:"sshHost,omitempty"`
	SshUser       string `json:"sshUser,omitempty"`
	SshPassword   string `json:"sshPassword,omitempty"`
	SshType       string `json:"sshType,omitempty"`
	SshKeyPath    string `json:"sshKeyPath,omitempty"`
	SshPort       string `json:"sshPort,omitempty"`
	AppType       string `json:"appType,omitempty"`
	Port          string `json:"port,omitempty"`
	Extra         string `json:"extra"`
	Agent         string `json:"agent"`
}
