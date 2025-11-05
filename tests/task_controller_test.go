package tests

import (
	"bytes"
	"encoding/json"
	"go-task-manager-mvc/config"
	"go-task-manager-mvc/models"
	"go-task-manager-mvc/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	rootPath, err := filepath.Abs("../")
	if err != nil {
		log.Fatal(err)
	}

	envPath := filepath.Join(rootPath, ".env")
	if _, err := os.Stat(envPath); err == nil {
		err = godotenv.Load(envPath)
		if err != nil {
			log.Printf("No se pudo cargar el .env desde %s : %v", envPath, err)
		} else {
			log.Printf(".env cargado desde %s", envPath)
		}
	} else {
		log.Printf("No se encontró el .env en %s", envPath)
	}
}

func TestCreateTask(t *testing.T) {
	config.ConnectDB()
	models.MigrateModels()
	r := gin.Default()
	routes.SetupRoutes(r)

	task := map[string]string{
		"title":       "Test Task",
		"description": "Testing creation",
		"status":      "pendiente",
		"due_date":    "2025-11-10T00:00:00Z",
	}

	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/task/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetTasks(t *testing.T) {
	config.ConnectDB()
	models.MigrateModels()
	r := gin.Default()
	routes.SetupRoutes(r)

	req, _ := http.NewRequest("GET", "/task/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
