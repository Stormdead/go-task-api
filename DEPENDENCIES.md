# 📦 Dependencias del Proyecto

Este archivo documenta todas las dependencias del proyecto Go Task Manager API.

## 🎯 Dependencias Principales

### Framework Web
- **gin-gonic/gin** `v1.10.0`
  - Framework HTTP de alto rendimiento
  - Router rápido y middleware
  - [Documentación](https://gin-gonic.com/)

### Base de Datos
- **gorm.io/gorm** `v1.25.12`
  - ORM para Go con soporte completo
  - Migraciones automáticas
  - [Documentación](https://gorm.io/)

- **gorm.io/driver/mysql** `v1.5.7`
  - Driver MySQL para GORM
  - Optimizado para MySQL 8.0+

### Autenticación y Seguridad
- **golang-jwt/jwt/v5** `v5.2.1`
  - Implementación de JSON Web Tokens
  - Manejo seguro de claims
  - [Documentación](https://github.com/golang-jwt/jwt)

- **golang.org/x/crypto** `v0.23.0`
  - Bcrypt para hash de contraseñas
  - Algoritmos criptográficos seguros

### Configuración
- **joho/godotenv** `v1.5.1`
  - Carga variables de entorno desde .env
  - [Documentación](https://github.com/joho/godotenv)

## 🧪 Dependencias de Testing

- **stretchr/testify** `v1.9.0`
  - Assertions y mocks para tests
  - Suite de testing completa
  - [Documentación](https://github.com/stretchr/testify)

## 📋 Instalación

### Opción 1: Automática
```bash
go mod download