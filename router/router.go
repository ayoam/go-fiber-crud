package router

import (
	"github.com/gofiber/fiber/v2"
	"notes-api/controller"
)

func NewRouter(noteController *controller.NoteController) *fiber.App {
	router := fiber.New()

	//test
	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcome to golang!",
		})
	})

	//notes router
	router.Route("/notes", func(router fiber.Router) {
		//create a new note
		router.Post("/", noteController.CreateNote)
	})

	return router
}