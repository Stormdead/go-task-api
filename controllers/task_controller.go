package controllers

import (
	"go-task-manager-mvc/config"
	"go-task-manager-mvc/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Devuelve todas las tareas
func GetTasks(c *gin.Context) {
	var task []models.Task
	config.DB.Find(&task)
	c.JSON(http.StatusOK, task)
}

// Filtra las tareas por estado
func GetTasksByStatus(c *gin.Context) {
	status := c.Query("status")

	var tasks []models.Task
	query := config.DB

	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Find(&tasks)
	c.JSON(200, tasks)
}

// Crea una nueva tarea
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&task)
	c.JSON(http.StatusCreated, task)

	if task.Title == "" {
		c.JSON(http.StatusFound, gin.H{"error": "El titulo es obligatorio"})
		return
	}

	validStatues := map[string]bool{
		"pendiente":   true,
		"en progreso": true,
		"completada":  true,
	}
	if !validStatues[task.Status] {
		c.JSON(http.StatusFound, gin.H{"Error": "Estado invalido. Tiene que usar Pendiente, en progreso o completada"})
		return
	}
}

// Actualiza una tarea existente
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
		return
	}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// Elimina una tarea
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
		return
	}
	config.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Tarea eliminada"})
}
