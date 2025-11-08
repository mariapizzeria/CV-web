package education

import "gorm.io/gorm"

type Education struct {
	gorm.Model
	Duration string `json:"duration"`
	College  string `json:"college"`
	Course   string `json:"course"`
}

type EducationResponse struct {
	Duration string `json:"duration"`
	College  string `json:"college"`
	Course   string `json:"course"`
}
