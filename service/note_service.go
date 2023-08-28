package service

import (
	"github.com/go-playground/validator/v10"
	"notes-api/converter"
	"notes-api/dto/request"
	"notes-api/dto/response"
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

func (serv *NoteService) CreateNote(noteDto request.CreateUpdateNoteRequestDto) (response.NoteResponseDto, error) {
	err := serv.Validate.Struct(noteDto)
	if err != nil {
		return response.NoteResponseDto{}, err
	}
	noteModel := serv.NoteConverter.DtoToNoteModel(noteDto)
	return serv.NoteConverter.NoteModelToDto(serv.NoteRepository.Save(noteModel)), nil
}

func (serv *NoteService) GetNoteById(noteId int) (response.NoteResponseDto, error) {
	noteModel, err := serv.NoteRepository.FindById(noteId)
	if err != nil {
		return response.NoteResponseDto{}, err
	}

	return serv.NoteConverter.NoteModelToDto(noteModel), nil
}

func (serv *NoteService) GetAllNotes() ([]response.NoteResponseDto, error) {
	noteModelList, err := serv.NoteRepository.FindAll()
	if err != nil {
		return []response.NoteResponseDto{}, err
	}

	var noteModelListResponse []response.NoteResponseDto
	for _, note := range noteModelList {
		noteModelListResponse = append(noteModelListResponse, serv.NoteConverter.NoteModelToDto(note))
	}

	return noteModelListResponse, nil
}

func (serv *NoteService) UpdateNote(noteId int, dto request.CreateUpdateNoteRequestDto) error {
	noteModel, err := serv.NoteRepository.FindById(noteId)
	if err != nil {
		return err
	}

	//update content
	noteModel.Content = dto.Content

	serv.NoteRepository.Save(noteModel)

	return nil
}

func (serv *NoteService) DeleteNote(noteId int) error {
	noteModel, err := serv.NoteRepository.FindById(noteId)
	if err != nil {
		return err
	}

	serv.NoteRepository.Delete(noteModel)

	return nil
}
