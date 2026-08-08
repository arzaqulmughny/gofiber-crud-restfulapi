package services

import (
	"errors"
	"gofiber-restful-api/database"
	"gofiber-restful-api/domain"
	"gofiber-restful-api/dtos"
	"gofiber-restful-api/repository"

	"github.com/gofiber/fiber/v3"
)

func CreateNotesService(notes *domain.Notes) (uint, error) {
	return repository.StoreNotes(notes)
}

func FindAllNotesSerivice() ([]domain.Notes, error) {
	return repository.FindAllNotes()
}

func FindNoteByIdService(id uint) (*domain.Notes, error) {
	return repository.FindNoteById(id)
}

func DeleteNoteByIdService(id uint) error {
	_, err := repository.FindNoteById(id)

	if err != nil {
		if err == fiber.ErrNotFound {
			return errors.New("record not found")
		}

		return err
	}

	return repository.DeleteNoteById(id)
}

func UpdateNoteByIdService(id uint, request *dtos.UpdateNoteRequest) (*domain.Notes, error) {
	// Check is note exist
	note, err := repository.FindNoteById(id)

	if err != nil {
		if err == fiber.ErrNotFound {
			return nil, errors.New("record not found")
		}

		return nil, err
	}

	// Update data
	note.Title = request.Title
	note.Body = request.Body
	database.DB.Save(&note)

	return note, nil
}
