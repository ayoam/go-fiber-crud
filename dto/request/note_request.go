package request

type CreateUpdateNoteRequest struct {
	Content string `json:"content" validate:"required,min=2,max=100"`
}
