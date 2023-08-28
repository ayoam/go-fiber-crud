package controller

import (
	"github.com/gofiber/fiber/v2"
	"notes-api/dto/request"
	"notes-api/dto/response"
	"notes-api/service"
	"strconv"
)

type NoteController struct {
	NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) *NoteController {
	return &NoteController{NoteService: noteService}
}

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

	webResponse := response.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "CREATED",
		Data:   noteResponseDto,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

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
	webResponse := response.WebResponse{
		Data: noteResponseDto,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (ctrl *NoteController) GetAllNotes(ctx *fiber.Ctx) error {
	notesListResponseDto, err := ctrl.NoteService.GetAllNotes()

	if err != nil {
		return err
	}
	webResponse := response.WebResponse{
		Data: notesListResponseDto,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

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

	err = ctrl.NoteService.UpdateNote(id, requestDto)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

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
