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
)

// TestUser representa un usuario de prueba
type TestUser struct {
	Username string
	Email    string
	Password string
	Token    string
	ID       uint
}

// init carga las variables de entorno
func init() {
	gin.SetMode(gin.TestMode)

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

// setupTestDB inicializa la base de datos de prueba
func setupTestDB() {
	config.ConnectDB()
	models.MigrateModels()
}

// setupRouter configura el router para tests
func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.SetupRoutes(r)
	return r
}

// cleanupTestData limpia los datos de prueba
func cleanupTestData() {
	config.DB.Exec("DELETE FROM tasks WHERE title LIKE '%TEST%'")
	config.DB.Exec("DELETE FROM users WHERE username LIKE '%testuser%'")
}

// createTestUser crea un usuario de prueba y retorna su token
func createTestUser(t *testing.T, router *gin.Engine, username string) TestUser {
	user := TestUser{
		Username: username,
		Email:    username + "@test.com",
		Password: "password123",
	}

	// Registrar usuario
	registerData := map[string]string{
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
	}

	body, _ := json.Marshal(registerData)
	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Logf("Failed to create user: %s", w.Body.String())
	}

	// Hacer login para obtener el token
	loginData := map[string]string{
		"email":    user.Email,
		"password": user.Password,
	}

	body, _ = json.Marshal(loginData)
	req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var loginResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &loginResponse)

	if token, ok := loginResponse["token"].(string); ok {
		user.Token = token
	}

	if userData, ok := loginResponse["user"].(map[string]interface{}); ok {
		if id, ok := userData["id"].(float64); ok {
			user.ID = uint(id)
		}
	}

	return user
}

// makeAuthenticatedRequest realiza una petición autenticada
func makeAuthenticatedRequest(method, url string, token string, body interface{}) (*httptest.ResponseRecorder, *http.Request) {
	var reqBody *bytes.Buffer
	if body != nil {
		jsonBody, _ := json.Marshal(body)
		reqBody = bytes.NewBuffer(jsonBody)
	} else {
		reqBody = bytes.NewBuffer([]byte{})
	}

	req, _ := http.NewRequest(method, url, reqBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	return w, req
}
