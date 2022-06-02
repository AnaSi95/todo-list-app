package controllers

import (
	"github.com/AnaSi95/todo-list-app/TodoLists"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateTaskInput struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskInput struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// GET /tasks
// Получаем список всех задач
func GetAllTasks(c *gin.Context) {
	var tasks []TodoLists.Item
	TodoLists.DB.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// GET /tasks/:id
// Получение одной задачи по ID
func GetTask(c *gin.Context) {
	// Проверяем имеется ли запись
	var task TodoLists.Item
	if err := TodoLists.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": task})
}

// POST /tasks
// Создание задачи
func CreateTask(c *gin.Context) {
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := TodoLists.Item{Title: input.Title, Description: input.Description}
	TodoLists.DB.Create(&task)

	c.JSON(http.StatusOK, gin.H{"tasks": task})
}

// PATCH /tasks/:id
// Изменения информации
func UpdateTask(c *gin.Context) {
	// Проверяем имеется ли такая запись перед тем как её менять
	var task TodoLists.Item
	if err := TodoLists.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	TodoLists.DB.Model(&task).Update(input)

	c.JSON(http.StatusOK, gin.H{"task": task})
}

// DELETE /tasks/:id
// Удаление
func DeleteTask(c *gin.Context) {
	// Проверяем имеется ли такая запись перед тем как её удалять
	var task TodoLists.Item
	if err := TodoLists.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	TodoLists.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"tasks": true})
}
