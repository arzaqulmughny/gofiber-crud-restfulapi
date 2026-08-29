package dtos

import "gofiber-restful-api/domain"

type CreateNewNoteRequest struct {
	Title string `json="title" validate:"required,min=3,max=255"`
	Body  string `json="body"`
}

type CreateNewNoteResponse struct {
	Error string       `json="error"`
	Data  domain.Notes `json="data"`
}

type UpdateNoteRequest struct {
	Title string `json:"title" validate:"required,min=3,max=255"`
	Body  string `json:"body"`
}

type FindAllNotesResponse struct {
	Error string         `json="error"`
	Data  []domain.Notes `json="data"`
}
