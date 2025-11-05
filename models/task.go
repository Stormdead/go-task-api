package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Status      string         `gorm:"default:'pendiente'" json:"status"`
	DueDate     *time.Time     `json:"due_date"`
	UserID      uint           `gorm:"not null;index" json:"user_id"`
	User        User           `gorm:"foreignKey:UserID" json:"-"` // json:"-" evita que se serialice en las respuestas
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeSave hook de GORM para validar antes de guardar
func (t *Task) BeforeSave(tx *gorm.DB) error {
	return t.Validate()
}

// BeforeUpdate hook de GORM para validar antes de actualizar
func (t *Task) BeforeUpdate(tx *gorm.DB) error {
	return t.Validate()
}

// Validate valida los campos de la tarea
func (t *Task) Validate() error {
	// Validar título
	t.Title = strings.TrimSpace(t.Title)
	if t.Title == "" {
		return errors.New("el título es obligatorio")
	}
	if len(t.Title) > TaskTitleMaxLength {
		return errors.New("el título no puede exceder los 200 caracteres")
	}

	// Validar descripción
	t.Description = strings.TrimSpace(t.Description)
	if len(t.Description) > TaskDescriptionMaxLength {
		return errors.New("la descripción no puede exceder los 1000 caracteres")
	}

	// Validar estado
	t.Status = strings.TrimSpace(strings.ToLower(t.Status))
	if t.Status == "" {
		t.Status = TaskStatusPending
	}
	if !IsValidTaskStatus(t.Status) {
		return errors.New("estado inválido. Use: pendiente, en progreso o completada")
	}

	// Validar fecha límite (no puede ser en el pasado para tareas nuevas)
	if t.DueDate != nil && t.ID == 0 { // Solo validar en creación
		if t.DueDate.Before(time.Now()) {
			return errors.New("la fecha límite no puede ser en el pasado")
		}
	}

	// Validar UserID
	if t.UserID == 0 {
		return errors.New("el usuario es obligatorio")
	}

	return nil
}

// IsOverdue verifica si la tarea está vencida
func (t *Task) IsOverdue() bool {
	if t.DueDate == nil {
		return false
	}
	return t.DueDate.Before(time.Now()) && t.Status != TaskStatusCompleted
}

// IsCompleted verifica si la tarea está completada
func (t *Task) IsCompleted() bool {
	return t.Status == TaskStatusCompleted
}

// MarkAsCompleted marca la tarea como completada
func (t *Task) MarkAsCompleted() {
	t.Status = TaskStatusCompleted
}

// GetStatusColor retorna un color para el estado (útil para frontend)
func (t *Task) GetStatusColor() string {
	switch t.Status {
	case TaskStatusPending:
		return "yellow"
	case TaskStatusInProgress:
		return "blue"
	case TaskStatusCompleted:
		return "green"
	default:
		return "gray"
	}
}
