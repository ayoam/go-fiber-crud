package converter

import (
	"notes-api/dto/request"
	"notes-api/dto/response"
	"notes-api/model"
)

type NoteConverter struct{}

func NewNoteConverter() *NoteConverter {
	return &NoteConverter{}
}

func (conv *NoteConverter) DtoToNoteModel(dto request.CreateUpdateNoteRequestDto) model.Note {
	noteModel := model.Note{
		Content: dto.Content,
	}
	return noteModel
}

func (conv *NoteConverter) NoteModelToDto(note model.Note) response.NoteResponseDto {
	dto := response.NoteResponseDto{
		ID:      note.ID,
		Content: note.Content,
	}
	return dto
}
