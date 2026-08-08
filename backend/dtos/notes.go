package dtos

type UpdateNoteRequest struct {
	Title string `json:"title" validate:"required,min=3,max=255"`
	Body  string `json:"body"`
}
