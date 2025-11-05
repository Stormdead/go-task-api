package main

import (
	"go-task-manager-mvc/config"
	"go-task-manager-mvc/models"
	"go-task-manager-mvc/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	log.Println("Conectado a la base de datos")
	models.MigrateModels()
	log.Println("Migraciones completadas")

	r := gin.Default()
	routes.SetupRoutes(r)
	log.Println("Servidor corriendo en http://localhost:8080")
	r.Run(":8080")
}
