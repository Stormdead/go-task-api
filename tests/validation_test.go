package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskValidations(t *testing.T) {
	setupTestDB()
	router := setupRouter()
	defer cleanupTestData()

	testUser := createTestUser(t, router, "testuser_validations")

	t.Run("Validar estados permitidos", func(t *testing.T) {
		validStatuses := []string{"pendiente", "en progreso", "completada"}

		for _, status := range validStatuses {
			taskData := map[string]interface{}{
				"title":  "TEST: Validación de estado " + status,
				"status": status,
			}

			w, req := makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusCreated, w.Code, "Estado '%s' debería ser válido", status)
		}
	})

	t.Run("Validar descripción con espacios", func(t *testing.T) {
		taskData := map[string]interface{}{
			"title":       "   TEST: Título con espacios   ",
			"description": "   Descripción con espacios   ",
		}

		w, req := makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		task := response["task"].(map[string]interface{})

		// Verificar que los espacios fueron eliminados
		assert.Equal(t, "TEST: Título con espacios", task["title"])
		assert.Equal(t, "Descripción con espacios", task["description"])
	})

	t.Run("Validar longitud máxima de descripción", func(t *testing.T) {
		longDescription := ""
		for i := 0; i < 1100; i++ {
			longDescription += "a"
		}

		taskData := map[string]interface{}{
			"title":       "TEST: Descripción muy larga",
			"description": longDescription,
		}

		w, req := makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Validar estado por defecto", func(t *testing.T) {
		taskData := map[string]interface{}{
			"title": "TEST: Sin estado definido",
		}

		w, req := makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		task := response["task"].(map[string]interface{})

		// El estado por defecto debería ser "pendiente"
		assert.Equal(t, "pendiente", task["status"])
	})
}
