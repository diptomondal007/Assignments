package crud

import "usercrudwithmocktest/crud/models"

type Repository interface {
	Get(page, limit int) ([]models.User, error)
	GetByUsername(username string) (*models.User, error)
	Create(user *models.User) error
	Delete(username string) (int64,error)
	Update(username string, user *models.User) error
}
