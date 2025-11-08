package experience

import "gorm.io/gorm"

type Experience struct {
	gorm.Model
	Duration    string `json:"duration"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ExperienceResponse struct {
	Duration    string `json:"duration"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
