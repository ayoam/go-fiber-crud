package controller

import (
	"github.com/gofiber/fiber/v2"
	"notes-api/dto/request"
	"notes-api/dto/response"
	"notes-api/helper"
	"notes-api/service"
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
	helper.ErrorPanic(err)

	noteResponseDto := ctrl.NoteService.CreateNote(requestDto)

	webResponse := response.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "CREATED",
		Data:   noteResponseDto,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
