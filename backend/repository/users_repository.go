package repository

import (
	"gofiber-restful-api/database"
	"gofiber-restful-api/domain"
)

func StoreUser(user domain.Users) (*domain.Users, error) {
	result := database.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
