package main

import (
	"github.com/AnaSi95/todo-list-app/TodoLists"
	"github.com/AnaSi95/todo-list-app/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	// Подключение к базе данных
	TodoLists.ConnectDB()

	// Маршруты
	route.GET("/tasks", controllers.GetAllTasks)
	route.POST("/tasks", controllers.CreateTask)
	route.GET("/tasks/:id", controllers.GetTask)
	route.PATCH("/tasks/:id", controllers.UpdateTask)
	route.DELETE("/tasks/:id", controllers.DeleteTask)

	// Запуск сервера
	route.Run()
}
