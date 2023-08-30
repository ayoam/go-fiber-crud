package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"notes-api/config"
	"notes-api/controller"
	"notes-api/converter"
	_ "notes-api/docs"
	"notes-api/model"
	"notes-api/repository"
	"notes-api/router"
	"notes-api/service"
)

// @title Go Fiber Rest API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /api/v1/
func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load env variables", err)
	}

	//database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	//Init repository
	noteRepository := repository.NewNoteRepository(db)

	//Init converter
	noteConverter := converter.NewNoteConverter()

	//Init service
	noteService := service.NewNoteService(*noteRepository, *noteConverter, validate)

	//Init controller
	noteController := controller.NewNoteController(*noteService)

	//Router
	routes := router.NewRouter(noteController)

	app := fiber.New()

	//Error middleware
	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": err.Error(),
			})
		}
		return nil
	})

	app.Mount("/api/v1", routes)
	log.Fatal(app.Listen(":8000"))
}
