package main

import (
	"fmt"
	"net/http"

	"github.com/anharathoi/go-rest-api-course/internal/database"
	transportHTTP "github.com/anharathoi/go-rest-api-course/internal/transport/http"
)

// App - the struct which contains things like pointer to db connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up app")

	var err error

	_, err = database.NewDatabase()

	if err != nil {
		return nil
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("failed to setup server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting the API")
		fmt.Println(err)
	}
}
