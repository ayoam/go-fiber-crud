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

func NewNoteRepository(Db *gorm.DB) *NoteRepository {
	return &NoteRepository{Db: Db}
}

func (rep *NoteRepository) Save(note model.Note) model.Note {
	var existingNote model.Note

	// Try to find an existing note by ID
	result := rep.Db.First(&existingNote, note.ID)
	if result.Error != nil {
		// User not found, create a new one
		result = rep.Db.Create(&note)
		return note
	}

	// Update the existing note's information
	existingNote.ID = note.ID
	existingNote.Content = note.Content

	// Save the updated note back to the database
	result = rep.Db.Save(&existingNote)
	return existingNote
}

func (rep *NoteRepository) Delete(note model.Note) {
	result := rep.Db.Delete(&note)
	helper.ErrorPanic(result.Error)
}

func (rep *NoteRepository) FindById(noteId int) (model.Note, error) {
	var note model.Note
	result := rep.Db.First(&note, noteId)
	if result.Error != nil {
		return note, errors.New("NOTE_NOT_FOUND")
	} else {
		return note, nil
	}
}

func (rep *NoteRepository) FindAll() ([]model.Note, error) {
	var notes []model.Note
	result := rep.Db.Find(&notes)
	if result.Error != nil {
		return nil, result.Error
	}
	return notes, nil
}
