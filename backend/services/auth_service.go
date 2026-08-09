package services

import (
	"errors"
	"gofiber-restful-api/database"
	"gofiber-restful-api/domain"
	"gofiber-restful-api/dtos"
	"gofiber-restful-api/repository"
	"gofiber-restful-api/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserLoginService(request dtos.UserLoginRequest) (string, error) {
	var user domain.Users

	// Search from users table, is user with username from request exists?
	result := database.DB.Where("username = ?", request.Username).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", errors.New("record not found")
		}

		return "", result.Error
	}

	// If exists, check is password valid
	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	)

	if err != nil {
		return "", errors.New("record not found")
	}

	token, err := utils.GenerateToken(user.Id)

	if err != nil {
		return "", err
	}

	return token, nil
}

func UserRegisterService(request dtos.UserRegisterRequest) (*domain.Users, error) {
	// Generate hashed password
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user, err := repository.StoreUser(domain.Users{
		Username: request.Username,
		Password: string(hash),
		Email:    request.Email,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
