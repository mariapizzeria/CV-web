package stack

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Stack struct {
	gorm.Model
	Type  string         `json:"type"`
	Skill pq.StringArray `json:"skill" gorm:"type:text[]"`
}

type StackResponse struct {
	Type  string         `json:"type"`
	Skill pq.StringArray `json:"skill" gorm:"type:text[]"`
}
