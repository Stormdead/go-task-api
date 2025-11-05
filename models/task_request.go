package models

import "time"

// representa los datos para crear una tarea
type TaskCreateRequest struct {
	Title       string     `json:"title" binding:"required,min=3,max=200"`
	Description string     `json:"description" binding:"max=1000"`
	Status      string     `json:"status" binding:"omitempty,oneof=pendiente 'en progreso' completada"`
	DueDate     *time.Time `json:"due_date"`
}

// representa los datos para actualizar una tarea
type TaskUpdateRequest struct {
	Title       string     `json:"title" binding:"omitempty,min=3,max=200"`
	Description string     `json:"description" binding:"omitempty,max=1000"`
	Status      string     `json:"status" binding:"omitempty,oneof=pendiente 'en progreso' completada"`
	DueDate     *time.Time `json:"due_date"`
}

// convierte TaskCreateRequest a Task
func (r *TaskCreateRequest) ToTask(userID uint) Task {
	status := r.Status
	if status == "" {
		status = TaskStatusPending
	}

	return Task{
		Title:       r.Title,
		Description: r.Description,
		Status:      status,
		DueDate:     r.DueDate,
		UserID:      userID,
	}
}

// aplica los cambios del TaskUpdateRequest a un Task existente
func (r *TaskUpdateRequest) ApplyToTask(task *Task) {
	if r.Title != "" {
		task.Title = r.Title
	}
	if r.Description != "" {
		task.Description = r.Description
	}
	if r.Status != "" {
		task.Status = r.Status
	}
	if r.DueDate != nil {
		task.DueDate = r.DueDate
	}
}
