package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	setupTestDB()
	router := setupRouter()
	defer cleanupTestData()

	testUser := createTestUser(t, router, "testuser_create_task")

	t.Run("Crear tarea exitosamente", func(t *testing.T) {
		taskData := map[string]interface{}{
			"title":       "TEST: Tarea de prueba",
			"description": "Descripción de prueba",
			"status":      "pendiente",
			"due_date":    "2025-12-31T23:59:59Z",
		}

		w, req := makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "Tarea creada exitosamente", response["message"])
		assert.NotNil(t, response["task"])
	})

	t.Run("Crear tarea sin autenticación", func(t *testing.T) {
		taskData := map[string]interface{}{
			"title": "TEST: Tarea sin auth",
		}

		w, req := makeAuthenticatedRequest("POST", "/api/tasks", "", taskData)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Crear tarea sin título", func(t *testing.T) {
		taskData := map[string]interface{}{
			"description": "Sin título",
		}

		w, req := makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Crear tarea con estado inválido", func(t *testing.T) {
		taskData := map[string]interface{}{
			"title":  "TEST: Tarea con estado inválido",
			"status": "estado_invalido",
		}

		w, req := makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Crear tarea con título muy largo", func(t *testing.T) {
		longTitle := ""
		for i := 0; i < 250; i++ {
			longTitle += "a"
		}

		taskData := map[string]interface{}{
			"title": longTitle,
		}

		w, req := makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Crear tarea con fecha pasada", func(t *testing.T) {
		taskData := map[string]interface{}{
			"title":    "TEST: Tarea con fecha pasada",
			"due_date": "2020-01-01T00:00:00Z",
		}

		w, req := makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestGetTasks(t *testing.T) {
	setupTestDB()
	router := setupRouter()
	defer cleanupTestData()

	testUser := createTestUser(t, router, "testuser_get_tasks")

	// Crear algunas tareas
	taskData := map[string]interface{}{
		"title":       "TEST: Tarea 1",
		"description": "Descripción 1",
		"status":      "pendiente",
	}
	w, req := makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
	router.ServeHTTP(w, req)

	taskData["title"] = "TEST: Tarea 2"
	taskData["status"] = "completada"
	w, req = makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
	router.ServeHTTP(w, req)

	t.Run("Obtener todas las tareas", func(t *testing.T) {
		w, req := makeAuthenticatedRequest("GET", "/api/tasks", testUser.Token, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.NotNil(t, response["tasks"])
		assert.NotZero(t, response["count"])
	})

	t.Run("Obtener tareas sin autenticación", func(t *testing.T) {
		w, req := makeAuthenticatedRequest("GET", "/api/tasks", "", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Usuario solo ve sus propias tareas", func(t *testing.T) {
		// Crear otro usuario
		testUser2 := createTestUser(t, router, "testuser_get_tasks_2")

		// Obtener tareas del usuario 2 (no debería ver las del usuario 1)
		w, req := makeAuthenticatedRequest("GET", "/api/tasks", testUser2.Token, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)

		// El usuario 2 no debería tener tareas (o muy pocas, no las del usuario 1)
		tasks := response["tasks"].([]interface{})
		for _, task := range tasks {
			taskMap := task.(map[string]interface{})
			// Verificar que el user_id no sea del usuario 1
			assert.NotEqual(t, float64(testUser.ID), taskMap["user_id"])
		}
	})
}

func TestUpdateTask(t *testing.T) {
	setupTestDB()
	router := setupRouter()
	defer cleanupTestData()

	testUser := createTestUser(t, router, "testuser_update_task")

	// Crear una tarea
	taskData := map[string]interface{}{
		"title":  "TEST: Tarea para actualizar",
		"status": "pendiente",
	}
	w, req := makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
	router.ServeHTTP(w, req)

	var createResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &createResponse)
	taskID := createResponse["task"].(map[string]interface{})["id"].(float64)

	t.Run("Actualizar tarea exitosamente", func(t *testing.T) {
		updateData := map[string]interface{}{
			"title":  "TEST: Tarea actualizada",
			"status": "completada",
		}

		w, req := makeAuthenticatedRequest("PUT", fmt.Sprintf("/api/tasks/%d", int(taskID)), testUser.Token, updateData)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "Tarea actualizada exitosamente", response["message"])
	})

	t.Run("Actualizar tarea inexistente", func(t *testing.T) {
		updateData := map[string]interface{}{
			"title": "TEST: Tarea que no existe",
		}

		w, req := makeAuthenticatedRequest("PUT", "/api/tasks/99999", testUser.Token, updateData)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Actualizar tarea de otro usuario", func(t *testing.T) {
		// Crear otro usuario
		testUser2 := createTestUser(t, router, "testuser_update_task_2")

		// Intentar actualizar la tarea del primer usuario
		updateData := map[string]interface{}{
			"title": "TEST: Intento de hackeo",
		}

		w, req := makeAuthenticatedRequest("PUT", fmt.Sprintf("/api/tasks/%d", int(taskID)), testUser2.Token, updateData)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestDeleteTask(t *testing.T) {
	setupTestDB()
	router := setupRouter()
	defer cleanupTestData()

	testUser := createTestUser(t, router, "testuser_delete_task")

	// Crear una tarea
	taskData := map[string]interface{}{
		"title": "TEST: Tarea para eliminar",
	}
	w, req := makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
	router.ServeHTTP(w, req)

	var createResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &createResponse)
	taskID := createResponse["task"].(map[string]interface{})["id"].(float64)

	t.Run("Eliminar tarea exitosamente", func(t *testing.T) {
		w, req := makeAuthenticatedRequest("DELETE", fmt.Sprintf("/api/tasks/%d", int(taskID)), testUser.Token, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "Tarea eliminada exitosamente", response["message"])
	})

	t.Run("Eliminar tarea inexistente", func(t *testing.T) {
		w, req := makeAuthenticatedRequest("DELETE", "/api/tasks/99999", testUser.Token, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Eliminar tarea de otro usuario", func(t *testing.T) {
		// Crear tarea con usuario 1
		taskData := map[string]interface{}{
			"title": "TEST: Tarea para proteger",
		}
		w, req := makeAuthenticatedRequest("POST", "/api/tasks", testUser.Token, taskData)
		router.ServeHTTP(w, req)

		var createResponse map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &createResponse)
		taskID := createResponse["task"].(map[string]interface{})["id"].(float64)

		// Crear otro usuario
		testUser2 := createTestUser(t, router, "testuser_delete_task_2")

		// Intentar eliminar con usuario 2
		w, req = makeAuthenticatedRequest("DELETE", fmt.Sprintf("/api/tasks/%d", int(taskID)), testUser2.Token, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
