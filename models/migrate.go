package models

import (
	"go-task-manager-mvc/config"
	"log"
)

func MigrateModels() {
	err := config.DB.AutoMigrate(&Task{}, &User{})
	if err != nil {
		log.Fatalf("Error al migrar modelos: %v", err)
	}
	log.Println("Migración de modelos completada exitosamente.")
}
