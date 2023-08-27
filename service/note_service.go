package service

import (
	"github.com/go-playground/validator/v10"
	"notes-api/converter"
	"notes-api/dto/request"
	"notes-api/dto/response"
	"notes-api/helper"
	"notes-api/repository"
)

type NoteService struct {
	NoteRepository repository.NoteRepository
	NoteConverter  converter.NoteConverter
	Validate       *validator.Validate
}

func NewNoteService(noteRepository repository.NoteRepository, noteConverter converter.NoteConverter, validate *validator.Validate) *NoteService {
	return &NoteService{
		NoteRepository: noteRepository,
		NoteConverter:  noteConverter,
		Validate:       validate,
	}
}

func (serv *NoteService) CreateNote(noteDto request.CreateUpdateNoteRequestDto) response.NoteResponseDto {
	err := serv.Validate.Struct(noteDto)
	helper.ErrorPanic(err)
	noteModel := serv.NoteConverter.DtoToNoteModel(noteDto)
	return serv.NoteConverter.NoteModelToDto(serv.NoteRepository.Save(noteModel))
}
