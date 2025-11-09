package education

import (
	"github.com/mariapizzeria/cv-web/backend/db"
	"gorm.io/gorm/clause"
)

type EducationRepository struct {
	db *db.Db
}

func NewRepository(database *db.Db) EducationRepository {
	return EducationRepository{database}
}

// Получение всех пройденных курсов
func (repo *EducationRepository) GetAll() ([]EducationResponse, error) {
	var result []EducationResponse
	res := repo.db.Table("education").Order("duration").Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	return result, nil
}

// Получение курса по id
func (repo *EducationRepository) GetById(id string) (*Education, error) {
	var result Education
	res := repo.db.Table("education").Where("id = ?", id).First(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	return &result, nil
}

// Обновление курса по id
func (repo *EducationRepository) UpdateContent(content *Education) (*Education, error) {
	res := repo.db.Table("education").Clauses(clause.Returning{}).Updates(content)
	if res.Error != nil {
		return nil, res.Error
	}
	return content, nil
}

// Добавление нового курса
func (repo *EducationRepository) CreateNewContent(content *EducationResponse) (*EducationResponse, error) {
	res := repo.db.Table("education").Create(content)
	if res.Error != nil {
		return nil, res.Error
	}
	return content, nil
}

// Удаление курса по id
func (repo *EducationRepository) DeleteContent(id uint) error {
	var education Education
	res := repo.db.Table("education").Delete(&education, id)
	return res.Error
}
