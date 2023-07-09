package service

import (
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/api/repository"
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/models"
)

// TextService TextService struct
type TextService struct {
	repository repository.TextRepository
}

// NewTextService : returns the TextService struct instance
func NewTextService(r repository.TextRepository) TextService {
	return TextService{
		repository: r,
	}
}

// Save -> calls text repository save method
func (repo TextService) Save(text models.Text) error {
	return repo.repository.Save(text)
}

// FindAll -> calls text repo find all method
func (repo TextService) FindAll(text models.Text, keyword string) (*[]models.Text, int64, error) {
	return repo.repository.FindAll(text, keyword)
}

// Update -> calls textrepo update method
func (repo TextService) Update(text models.Text) error {
	return repo.repository.Update(text)
}

// Delete -> calls text repo delete method
func (repo TextService) Delete(id int64) error {
	var text models.Text
	text.ID = id
	return repo.repository.Delete(text)
}

// Find -> calls text repo find method
func (repo TextService) Find(text models.Text) (models.Text, error) {
	return repo.repository.Find(text)
}
