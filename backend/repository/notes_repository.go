package repository

import (
	"gofiber-restful-api/database"
	"gofiber-restful-api/domain"
)

func StoreNotes(notes *domain.Notes) (uint, error) {
	result := database.DB.Create(notes)

	if (result.Error != nil) {
		return 0, result.Error
	}

	return notes.ID, nil
}

func FindAllNotes() ([]domain.Notes, error) {
	var notes []domain.Notes

	result := database.DB.Find(&notes)

	if (result.Error != nil) {
		return nil, result.Error
	}

	return notes, nil
}