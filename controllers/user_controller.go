package controllers

import (
	"net/http"

	"go-task-manager-mvc/config"
	"go-task-manager-mvc/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Registro de usuario
func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Encriptar contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encriptar la contraseña"})
		return
	}
	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo registrar el usuario"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado correctamente",
	})
}

// Login de usuario
func LoginUser(c *gin.Context) {
	var request models.User
	var user models.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Buscar usuario por email
	if err := config.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Comparar contraseñas
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
		return
	}

	// Generar token JWT
	token, err := config.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Inicio de sesión exitoso",
		"token":   token,
	})
}
