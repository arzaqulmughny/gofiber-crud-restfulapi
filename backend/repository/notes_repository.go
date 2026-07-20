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

func FindNoteById(id uint) (*domain.Notes, error) {
	var note domain.Notes

	result := database.DB.First(&note, id)
	if (result.Error != nil) {
		return nil, result.Error
	}

	return &note, nil
}

func DeleteNoteById(id uint) error {
	result := database.DB.Delete(&domain.Notes{}, id)
	
	if (result.Error != nil) {
		return result.Error
	}

	return nil
}