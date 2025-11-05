package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRegistration(t *testing.T) {
	setupTestDB()
	router := setupRouter()
	defer cleanupTestData()

	t.Run("Registro exitoso", func(t *testing.T) {
		userData := map[string]string{
			"username": "testuser_register",
			"email":    "testuser_register@test.com",
			"password": "password123",
		}

		body, _ := json.Marshal(userData)
		req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "Usuario registrado correctamente", response["message"])
		assert.NotNil(t, response["user"])
	})

	t.Run("Registro con email duplicado", func(t *testing.T) {
		// Crear primer usuario
		createTestUser(t, router, "testuser_duplicate")

		// Intentar crear usuario con mismo email
		userData := map[string]string{
			"username": "testuser_duplicate2",
			"email":    "testuser_duplicate@test.com",
			"password": "password123",
		}

		body, _ := json.Marshal(userData)
		req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Registro con datos inválidos", func(t *testing.T) {
		userData := map[string]string{
			"username": "ab", // Muy corto
			"email":    "invalid-email",
			"password": "123", // Muy corta
		}

		body, _ := json.Marshal(userData)
		req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUserLogin(t *testing.T) {
	setupTestDB()
	router := setupRouter()
	defer cleanupTestData()

	// Crear usuario de prueba
	testUser := createTestUser(t, router, "testuser_login")

	t.Run("Login exitoso", func(t *testing.T) {
		loginData := map[string]string{
			"email":    testUser.Email,
			"password": testUser.Password,
		}

		body, _ := json.Marshal(loginData)
		req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "Inicio de sesión exitoso", response["message"])
		assert.NotEmpty(t, response["token"])
		assert.NotNil(t, response["user"])
	})

	t.Run("Login con email inexistente", func(t *testing.T) {
		loginData := map[string]string{
			"email":    "noexiste@test.com",
			"password": "password123",
		}

		body, _ := json.Marshal(loginData)
		req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Login con contraseña incorrecta", func(t *testing.T) {
		loginData := map[string]string{
			"email":    testUser.Email,
			"password": "wrongpassword",
		}

		body, _ := json.Marshal(loginData)
		req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Login sin credenciales", func(t *testing.T) {
		loginData := map[string]string{}

		body, _ := json.Marshal(loginData)
		req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
