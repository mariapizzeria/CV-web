package experience

import (
	"cv-web/backend/pkg/db"

	"gorm.io/gorm/clause"
)

type ExpRepository struct {
	db *db.Db
}

func NewRepository(database *db.Db) *ExpRepository {
	return &ExpRepository{database}
}

func (repo *ExpRepository) GetAll() ([]ExperienceResponse, error) {
	var result []ExperienceResponse
	res := repo.db.Table("experience").Where("deleted_at is null").Order("duration DESC").Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	return result, nil
}

func (repo *ExpRepository) CreateNewContent(content *ExperienceResponse) (*ExperienceResponse, error) {
	res := repo.db.Table("experience").Create(&content)
	if res.Error != nil {
		return nil, res.Error
	}
	return content, nil
}

func (repo *ExpRepository) UpdateContent(content *Experience) (*Experience, error) {
	res := repo.db.Table("experience").Clauses(clause.Returning{}).Updates(content)
	if res.Error != nil {
		return nil, res.Error
	}
	return content, nil
}

func (repo *ExpRepository) DeleteContent(id uint) error {
	var experience Experience
	res := repo.db.Table("experience").Delete(&experience, id)
	return res.Error
}
