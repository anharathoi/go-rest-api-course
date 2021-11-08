package main

import "fmt"

// App - the struct which contains things like pointer to db connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up app")
	return nil
}

func main() {
	fmt.Println("Go REST API")
	app := App{}
	if err := app.Run; err != nil {
		fmt.Println("Error starting the API")
		fmt.Println(err)
	}
}
