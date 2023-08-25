package repository

import (
	"errors"
	"notes-api/helper"
	"notes-api/model"

	"gorm.io/gorm"
)

type NoteRepository struct {
	Db *gorm.DB
}

func NewNoteRespository(Db *gorm.DB) *NoteRepository {
	return &NoteRepository{Db: Db}
}

func (rep *NoteRepository) Save(note model.Note) model.Note {
	result := rep.Db.Create(&note)
	helper.ErrorPanic(result.Error)
	return note
}

func (rep *NoteRepository) Delete(note model.Note) {
	result := rep.Db.Delete(&note)
	helper.ErrorPanic(result.Error)
}

func (rep *NoteRepository) FindById(noteId int) (model.Note, error) {
	var note model.Note
	result := rep.Db.Find(&note, noteId)
	if result.Error != nil {
		return note, nil
	} else {
		return note, errors.New("note is not found")
	}
}

func (rep *NoteRepository) FindAll() []model.Note {
	var notes []model.Note
	result := rep.Db.Find(&notes)
	helper.ErrorPanic(result.Error)
	return notes
}
