package services

import (
	"errors"
	"ginTemplate/pkg/api/models"
)

var mockUsers = []models.User{
	{ID: "1", Name: "John Doe", Age: 30},
	{ID: "2", Name: "Jane Doe", Age: 25},
}

func GetUserByID(id string) (*models.User, error) {
	for _, user := range mockUsers {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
