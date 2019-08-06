package services

import (
	"test/database"
	"test/models"
)

type UserService struct {
}

func (s *UserService) GetById(id string) (*models.User, error) {
	user := &models.User{}
	if err := database.DB.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

func (s *UserService) Find() ([]*models.User, error) {
	users := make([]*models.User, 0)

	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
