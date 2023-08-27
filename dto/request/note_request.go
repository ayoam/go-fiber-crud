package request

type CreateUpdateNoteRequestDto struct {
	Content string `json:"content" validate:"required,min=2,max=100"`
}
