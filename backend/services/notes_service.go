package services

import (
	"gofiber-restful-api/domain"
	"gofiber-restful-api/repository"
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
	return repository.DeleteNoteById(id)
}