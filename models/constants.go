package models

//estados validos para las tareas
const (
	TaskStatusPending    = "pendiente"
	TaskStatusInProgress = "en progreso"
	TaskStatusCompleted  = "completada"
)

//Validaciones de longitud
const (
	TaskTitleMaxLength       = 200
	TaskDescriptionMaxLength = 1000
)

//Retorna el mapa de estados validos
func ValidTaskStatuses() map[string]bool {
	return map[string]bool{
		TaskStatusPending:    true,
		TaskStatusInProgress: true,
		TaskStatusCompleted:  true,
	}
}

//Verifica si un estado es valido
func IsValidTaskStatus(status string) bool {
	return ValidTaskStatuses()[status]
}

//Retorna una lista de estados validos
func GetValidTasksStatuesList() []string {
	return []string{
		TaskStatusPending,
		TaskStatusInProgress,
		TaskStatusCompleted,
	}
}
