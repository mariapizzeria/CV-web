package social

import (
	"cv-web/pkg/db"

	"gorm.io/gorm/clause"
)

type Repository struct {
	db *db.Db
}

func NewRepository(database *db.Db) *Repository {
	return &Repository{database}
}

func (repo *Repository) GetAll() ([]SocialResponse, error) {
	var result []SocialResponse
	res := repo.db.Table("social").Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	return result, nil
}

func (repo *Repository) CreateNewContent(content *SocialResponse) (*SocialResponse, error) {
	res := repo.db.Table("social").Create(&content)
	if res.Error != nil {
		return nil, res.Error
	}
	return content, nil
}

func (repo *Repository) UpdateContent(content *Social) (*Social, error) {
	res := repo.db.Table("social").Clauses(clause.Returning{}).Updates(&content)
	if res.Error != nil {
		return nil, res.Error
	}
	return content, nil
}

func (repo *Repository) DeleteContent(id uint) error {
	var social Social
	res := repo.db.Table("experience").Delete(&social, id)
	return res.Error
}
