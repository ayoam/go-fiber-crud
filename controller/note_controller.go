package controller

import (
	"github.com/gofiber/fiber/v2"
	"notes-api/dto/request"
	"notes-api/service"
	"strconv"
)

type NoteController struct {
	NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) *NoteController {
	return &NoteController{NoteService: noteService}
}

// CreateNote godoc
// @Summary Returns a 200
// @Param request body request.CreateUpdateNoteRequestDto true "query params"
// @Accept json
// @Produce json
// @Success 201 {object} model.Note
// @Failure 400
// @Router /notes [post]
func (ctrl *NoteController) CreateNote(ctx *fiber.Ctx) error {
	requestDto := request.CreateUpdateNoteRequestDto{}
	err := ctx.BodyParser(&requestDto)
	if err != nil {
		return err
	}

	noteResponseDto, err := ctrl.NoteService.CreateNote(requestDto)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(noteResponseDto)
}

// GetNoteById godoc
// @Summary Returns a 200
// @Param id path int true "Note ID"
// @Accept json
// @Produce json
// @Success 200 {object} model.Note
// @Failure 400
// @Router /notes/{id} [get]
func (ctrl *NoteController) GetNoteById(ctx *fiber.Ctx) error {
	noteId := ctx.Params("id")
	id, err := strconv.Atoi(noteId)
	if err != nil {
		return err
	}

	noteResponseDto, err := ctrl.NoteService.GetNoteById(id)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(noteResponseDto)
}

// GetAllNotes godoc
// @Summary Returns a 200
// @Accept json
// @Produce json
// @Success 200 {array} model.Note
// @Failure 400
// @Router /notes [get]
func (ctrl *NoteController) GetAllNotes(ctx *fiber.Ctx) error {
	notesListResponseDto, err := ctrl.NoteService.GetAllNotes()

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(notesListResponseDto)
}

// UpdateNote godoc
// @Summary Returns a 200
// @Param id path int true "Note ID"
// @Accept json
// @Produce json
// @Success 200 {array} model.Note
// @Failure 400
// @Router /notes [put]
func (ctrl *NoteController) UpdateNote(ctx *fiber.Ctx) error {
	noteId := ctx.Params("id")
	id, err := strconv.Atoi(noteId)
	if err != nil {
		return err
	}

	requestDto := request.CreateUpdateNoteRequestDto{}
	err = ctx.BodyParser(&requestDto)
	if err != nil {
		return err
	}

	responseDto, err := ctrl.NoteService.UpdateNote(id, requestDto)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusNoContent).JSON(responseDto)
}

// DeleteNote godoc
// @Summary Returns a 204
// @Param id path int true "Note ID"
// @Accept json
// @Success 204
// @Failure 400
// @Router /notes/{id} [delete]
func (ctrl *NoteController) DeleteNote(ctx *fiber.Ctx) error {
	noteId := ctx.Params("id")
	id, err := strconv.Atoi(noteId)
	if err != nil {
		return err
	}

	err = ctrl.NoteService.DeleteNote(id)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
