package stack

import "gorm.io/gorm"

type Stack struct {
	gorm.Model
	Type  string   `json:"type"`
	Skill []string `json:"skill"`
}

type StackResponse struct {
	Type  string   `json:"type"`
	Skill []string `json:"skill"`
}
