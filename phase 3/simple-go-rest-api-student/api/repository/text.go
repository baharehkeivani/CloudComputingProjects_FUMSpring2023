package repository

import (
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/infrastructure"
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/models"
)

// TextRepository -> TextRepository
type TextRepository struct {
	db infrastructure.Database
}

// NewTextRepository : fetching database
func NewTextRepository(db infrastructure.Database) TextRepository {
	return TextRepository{
		db: db,
	}
}

// Save -> Method for saving text to database
func (p TextRepository) Save(text models.Text) error {
	return p.db.DB.Create(&text).Error
}

// FindAll -> Method for fetching all texts from database
func (repo TextRepository) FindAll(text models.Text, keyword string) (*[]models.Text, int64, error) {
	var texts []models.Text
	var totalRows int64 = 0

	queryBuider := repo.db.DB.Order("created_at desc").Model(&models.Text{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			repo.db.DB.Where("text.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(text).
		Find(&texts).
		Count(&totalRows).Error
	return &texts, totalRows, err
}

// Update -> Method for updating Text
func (repo TextRepository) Update(text models.Text) error {
	return repo.db.DB.Save(&text).Error
}

// Find -> Method for fetching text by id
func (repo TextRepository) Find(text models.Text) (models.Text, error) {
	var texts models.Text
	err := repo.db.DB.
		Debug().
		Model(&models.Text{}).
		Where(&text).
		Take(&texts).Error
	return texts, err
}

// Delete Deletes text
func (repo TextRepository) Delete(text models.Text) error {
	return repo.db.DB.Delete(&text).Error
}
