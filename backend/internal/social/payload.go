package social

import "gorm.io/gorm"

type Social struct {
	gorm.Model
	Image string `json:"image"`
	Link  string `json:"link"`
}
type SocialResponse struct {
	Image string `json:"image"`
	Link  string `json:"link"`
}
