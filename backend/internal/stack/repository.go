package stack

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

func (repo *Repository) GetAll() ([]StackResponse, error) {
	var result []StackResponse
	res := repo.db.Table("skills").Where("deleted_at is NULL").Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	return result, nil
}

func (repo *Repository) CreateNewContent(content *StackResponse) (*StackResponse, error) {
	res := repo.db.Table("skills").Create(content)
	if res.Error != nil {
		return nil, res.Error
	}
	return content, nil
}

func (repo *Repository) UpdateContent(content *Stack) (*Stack, error) {
	res := repo.db.Table("skills").Clauses(clause.Returning{}).Updates(content)
	if res.Error != nil {
		return nil, res.Error
	}
	return content, nil
}

func (repo *Repository) DeleteContent(id uint) error {
	var result Stack
	res := repo.db.Table("skills").Delete(&result, id)
	return res.Error
}
