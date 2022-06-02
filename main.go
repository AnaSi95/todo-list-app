package main

import (
	todo_list_app "github.com/AnaSi95/todo-list-app"
	"log"
)

func main() {
	srv := new(todo_list_app.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
