package crud

import "usercrudwithmocktest/crud/models"

type crudService struct {
	crudRepo Repository
}

func (c crudService) Get(page, limit int) ([]models.User, error) {
	return c.crudRepo.Get(page,limit)
}

func (c crudService) GetByUsername(username string) (*models.User, error) {
	return c.crudRepo.GetByUsername(username)
}

func (c crudService) Create(user *models.User) error {
	return c.crudRepo.Create(user)
}

func (c crudService) Delete(username string) (int64,error ){
	return c.crudRepo.Delete(username)
}

func (c crudService) Update(username string, user *models.User) error {
	return c.crudRepo.Update(username, user)
}

func NewCrudService(r Repository) Service{
	return &crudService{crudRepo:r}
}