package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"notes-api/config"
	"notes-api/controller"
	"notes-api/converter"
	"notes-api/model"
	"notes-api/repository"
	"notes-api/router"
	"notes-api/service"
)

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

	app.Mount("/api/v1", routes)
	log.Fatal(app.Listen(":8000"))
}
