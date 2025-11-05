package controllers

import (
	"go-task-manager-mvc/config"
	"go-task-manager-mvc/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Helper function para obtener el UserID del contexto
func getUserIDFromContext(c *gin.Context) (uint, error) {
	username, exists := c.Get("username")
	if !exists {
		return 0, nil
	}

	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}

// GetTasks devuelve todas las tareas del usuario autenticado
func GetTasks(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autorizado"})
		return
	}

	var tasks []models.Task
	config.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&tasks)

	// Agregar información adicional
	tasksWithInfo := make([]gin.H, len(tasks))
	for i, task := range tasks {
		tasksWithInfo[i] = gin.H{
			"id":           task.ID,
			"title":        task.Title,
			"description":  task.Description,
			"status":       task.Status,
			"status_color": task.GetStatusColor(),
			"due_date":     task.DueDate,
			"is_overdue":   task.IsOverdue(),
			"is_completed": task.IsCompleted(),
			"user_id":      task.UserID,
			"created_at":   task.CreatedAt,
			"updated_at":   task.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasksWithInfo,
		"count": len(tasks),
	})
}

// GetTasksByStatus filtra las tareas por estado
func GetTasksByStatus(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autorizado"})
		return
	}

	status := c.Query("status")

	// Validar el estado si se proporciona
	if status != "" && !models.IsValidTaskStatus(status) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":          "Estado inválido",
			"valid_statuses": models.GetValidTasksStatuesList(),
		})
		return
	}

	var tasks []models.Task
	query := config.DB.Where("user_id = ?", userID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Order("created_at DESC").Find(&tasks)
	c.JSON(http.StatusOK, gin.H{
		"tasks":  tasks,
		"count":  len(tasks),
		"status": status,
	})
}

// CreateTask crea una nueva tarea
func CreateTask(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autorizado"})
		return
	}

	var request models.TaskCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Convertir request a Task
	task := request.ToTask(userID)

	// Crear la tarea (las validaciones se ejecutan en BeforeSave)
	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tarea creada exitosamente",
		"task": gin.H{
			"id":           task.ID,
			"title":        task.Title,
			"description":  task.Description,
			"status":       task.Status,
			"status_color": task.GetStatusColor(),
			"due_date":     task.DueDate,
			"is_overdue":   task.IsOverdue(),
			"user_id":      task.UserID,
			"created_at":   task.CreatedAt,
			"updated_at":   task.UpdatedAt,
		},
	})
}

// UpdateTask actualiza una tarea existente
func UpdateTask(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autorizado"})
		return
	}

	id := c.Param("id")
	var task models.Task

	// Buscar la tarea y verificar que pertenezca al usuario
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada o no tienes permiso para modificarla"})
		return
	}

	var request models.TaskUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Aplicar cambios
	request.ApplyToTask(&task)

	// Guardar cambios (las validaciones se ejecutan en BeforeUpdate)
	if err := config.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tarea actualizada exitosamente",
		"task": gin.H{
			"id":           task.ID,
			"title":        task.Title,
			"description":  task.Description,
			"status":       task.Status,
			"status_color": task.GetStatusColor(),
			"due_date":     task.DueDate,
			"is_overdue":   task.IsOverdue(),
			"is_completed": task.IsCompleted(),
			"user_id":      task.UserID,
			"updated_at":   task.UpdatedAt,
		},
	})
}

// DeleteTask elimina una tarea
func DeleteTask(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autorizado"})
		return
	}

	id := c.Param("id")
	var task models.Task

	// Buscar la tarea y verificar que pertenezca al usuario
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada o no tienes permiso para eliminarla"})
		return
	}

	// Eliminar la tarea (soft delete)
	if err := config.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la tarea"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tarea eliminada exitosamente",
		"task_id": id,
		"title":   task.Title,
	})
}
