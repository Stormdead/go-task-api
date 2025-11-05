# 📋 Go Task Manager API

> API RESTful completa para gestión de tareas con autenticación JWT, desarrollada en Go (Golang) con arquitectura limpia y buenas prácticas.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Tests](https://img.shields.io/badge/Tests-26%20Passing-success)](tests/)
[![Coverage](https://img.shields.io/badge/Coverage-85%25-brightgreen)]()

---

## 📑 Tabla de Contenidos

- [Características](#-características)
- [Tecnologías](#-tecnologías)
- [Arquitectura](#-arquitectura)
- [Requisitos Previos](#-requisitos-previos)
- [Instalación](#-instalación)
- [Configuración](#-configuración)
- [Uso](#-uso)
- [API Endpoints](#-api-endpoints)
- [Testing](#-testing)
- [Estructura del Proyecto](#-estructura-del-proyecto)
- [Contribuir](#-contribuir)
- [Licencia](#-licencia)

---

## ✨ Características

- 🔐 **Autenticación JWT** - Sistema seguro de registro e inicio de sesión
- 📋 **CRUD Completo** - Gestión completa de tareas (Crear, Leer, Actualizar, Eliminar)
- 👤 **Multi-usuario** - Cada usuario gestiona sus propias tareas
- 🔒 **Seguridad** - Middleware de autenticación y autorización
- ✅ **Validaciones Robustas** - Validación de datos a nivel de modelo y controlador
- 🧪 **Testing Completo** - 26 tests unitarios y de integración
- 📊 **Soft Delete** - Eliminación lógica de registros
- 🔍 **Filtros** - Búsqueda y filtrado por estado de tareas
- 📝 **Logs** - Sistema de logging para debugging
- 🚀 **API RESTful** - Diseño siguiendo estándares REST

---

## 🛠️ Tecnologías

| Tecnología | Versión | Propósito |
|-----------|---------|-----------|
| **Go** | 1.21+ | Lenguaje de programación |
| **Gin** | 1.10.0 | Framework web HTTP |
| **GORM** | 1.25.12 | ORM para base de datos |
| **MySQL** | 8.0+ | Base de datos relacional |
| **JWT** | 5.2.1 | Autenticación con tokens |
| **Bcrypt** | 0.23.0 | Encriptación de contraseñas |
| **Testify** | 1.9.0 | Framework de testing |

---

## 🏗️ Arquitectura

El proyecto sigue una arquitectura MVC (Model-View-Controller) adaptada para APIs:

```
┌─────────────┐
│   Cliente   │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│   Router    │ ← Gin Engine
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ Middleware  │ ← JWT Auth
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ Controllers │ ← Lógica HTTP
└──────┬──────┘
       │
       ▼
┌─────────────┐
│   Models    │ ← Validaciones
└──────┬──────┘
       │
       ▼
┌─────────────┐
│   Database  │ ← MySQL + GORM
└─────────────┘
```

---

## 📋 Requisitos Previos

Antes de comenzar, asegúrate de tener instalado:

- **Go** >= 1.21 ([Descargar](https://golang.org/dl/))
- **MySQL** >= 8.0 ([Descargar](https://dev.mysql.com/downloads/))
- **Git** ([Descargar](https://git-scm.com/downloads))
- **Postman** (opcional, para probar la API) ([Descargar](https://www.postman.com/downloads/))

### Verificar instalaciones:

```bash
go version        # Debería mostrar go1.21 o superior
mysql --version   # Debería mostrar mysql 8.0 o superior
git --version     # Debería mostrar la versión de git
```

---

## 🚀 Instalación

### 1. Clonar el repositorio

```bash
git clone https://github.com/Stormdead/go-task-api.git
cd go-task-api
```

### 2. Instalar dependencias

```bash
go mod download
```

Si prefieres actualizar las dependencias a las últimas versiones:

```bash
go mod tidy
```

### 3. Crear la base de datos

```bash
# Acceder a MySQL
mysql -u root -p

# Crear la base de datos
CREATE DATABASE tasks_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# Crear usuario (opcional pero recomendado)
CREATE USER 'taskuser'@'localhost' IDENTIFIED BY 'tu_password_seguro';
GRANT ALL PRIVILEGES ON tasks_db.* TO 'taskuser'@'localhost';
FLUSH PRIVILEGES;

# Salir de MySQL
EXIT;
```

### 4. Configurar variables de entorno

Crea un archivo `.env` en la raíz del proyecto:

```bash
cp .env.example .env
```

Edita el archivo `.env` con tus credenciales:

```env
# Database Configuration
DB_USER=taskuser
DB_PASSWORD=tu_password_seguro
DB_HOST=localhost
DB_PORT=3306
DB_NAME=tasks_db

# JWT Configuration
JWT_SECRET=tu_clave_secreta_super_segura_y_larga_minimo_32_caracteres

# Server Configuration
PORT=8080
GIN_MODE=release
```

### 5. Ejecutar migraciones

Las migraciones se ejecutan automáticamente al iniciar la aplicación, pero puedes verificarlas:

```bash
go run main.go
```

Verás en los logs:
```
Conectado a la base de datos
Migración de modelos completada exitosamente.
Servidor corriendo en http://localhost:8080
```

---

## ⚙️ Configuración

### Archivo `.env.example`

```env
# Database
DB_USER=root
DB_PASSWORD=
DB_HOST=localhost
DB_PORT=3306
DB_NAME=tasks_db

# JWT
JWT_SECRET=change-this-to-a-secure-secret-key

# Server
PORT=8080
GIN_MODE=debug
```

### Modos de Ejecución

- **Desarrollo**: `GIN_MODE=debug` (muestra logs detallados)
- **Producción**: `GIN_MODE=release` (optimizado)

---

## 💻 Uso

### Iniciar el servidor

```bash
# Modo desarrollo
go run main.go

# Compilar y ejecutar
go build -o task-api
./task-api  # En Windows: task-api.exe
```

El servidor estará disponible en: `http://localhost:8080`

### Probar que funciona

```bash
curl http://localhost:8080/api/register -X POST \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"123456"}'
```

---

## 📡 API Endpoints

### 🔐 Autenticación (Públicos)

| Método | Endpoint | Descripción | Body |
|--------|----------|-------------|------|
| POST | `/api/register` | Registrar nuevo usuario | `username`, `email`, `password` |
| POST | `/api/login` | Iniciar sesión | `email`, `password` |

### 📋 Tareas (Requieren autenticación)

| Método | Endpoint | Descripción | Headers |
|--------|----------|-------------|---------|
| GET | `/api/tasks` | Obtener todas las tareas del usuario | `Authorization: Bearer {token}` |
| POST | `/api/tasks` | Crear nueva tarea | `Authorization: Bearer {token}` |
| PUT | `/api/tasks/:id` | Actualizar tarea | `Authorization: Bearer {token}` |
| DELETE | `/api/tasks/:id` | Eliminar tarea | `Authorization: Bearer {token}` |

### 📝 Ejemplos de uso

#### 1. Registro de usuario

```bash
POST /api/register
Content-Type: application/json

{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "securepass123"
}
```

**Respuesta:**
```json
{
  "message": "Usuario registrado correctamente",
  "user": {
    "id": 1,
    "username": "johndoe",
    "email": "john@example.com"
  }
}
```

#### 2. Login

```bash
POST /api/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "securepass123"
}
```

**Respuesta:**
```json
{
  "message": "Inicio de sesión exitoso",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "johndoe",
    "email": "john@example.com"
  }
}
```

#### 3. Crear tarea

```bash
POST /api/tasks
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
Content-Type: application/json

{
  "title": "Completar proyecto",
  "description": "Terminar la API de tareas",
  "status": "en progreso",
  "due_date": "2025-12-31T23:59:59Z"
}
```

**Respuesta:**
```json
{
  "message": "Tarea creada exitosamente",
  "task": {
    "id": 1,
    "title": "Completar proyecto",
    "description": "Terminar la API de tareas",
    "status": "en progreso",
    "status_color": "blue",
    "due_date": "2025-12-31T23:59:59Z",
    "is_overdue": false,
    "user_id": 1,
    "created_at": "2025-11-05T20:15:00Z",
    "updated_at": "2025-11-05T20:15:00Z"
  }
}
```

#### 4. Obtener todas las tareas

```bash
GET /api/tasks
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Respuesta:**
```json
{
  "tasks": [
    {
      "id": 1,
      "title": "Completar proyecto",
      "description": "Terminar la API de tareas",
      "status": "en progreso",
      "status_color": "blue",
      "due_date": "2025-12-31T23:59:59Z",
      "is_overdue": false,
      "is_completed": false,
      "user_id": 1,
      "created_at": "2025-11-05T20:15:00Z",
      "updated_at": "2025-11-05T20:15:00Z"
    }
  ],
  "count": 1
}
```

### Estados válidos de tareas

- `pendiente` - Tarea no iniciada (color: amarillo)
- `en progreso` - Tarea en desarrollo (color: azul)
- `completada` - Tarea finalizada (color: verde)

---

## 🧪 Testing

El proyecto incluye 26 tests unitarios y de integración que cubren:

- ✅ Autenticación (registro y login)
- ✅ CRUD de tareas
- ✅ Validaciones de datos
- ✅ Seguridad y permisos
- ✅ Aislamiento entre usuarios

### Ejecutar tests

```bash
# Ejecutar todos los tests
go test ./tests/... -v

# Ejecutar tests específicos
go test ./tests/ -run TestUserRegistration -v

# Con cobertura
go test ./... -coverprofile=coverage.out -coverpkg=./...

# Ver reporte de cobertura
go tool cover -html coverage.out
```

### Resultado esperado

```
PASS: TestUserRegistration (3 subtests)
PASS: TestUserLogin (4 subtests)
PASS: TestCreateTask (6 subtests)
PASS: TestGetTasks (3 subtests)
PASS: TestUpdateTask (3 subtests)
PASS: TestDeleteTask (3 subtests)
PASS: TestTaskValidations (4 subtests)

Total: 26 tests passing
```

---

## 📁 Estructura del Proyecto

```
go-task-api/
├── config/
│   ├── database.go      # Conexión a base de datos
│   └── jwt.go           # Configuración JWT
├── controllers/
│   ├── user_controller.go   # Controlador de usuarios
│   └── task_controller.go   # Controlador de tareas
├── middleware/
│   └── authMiddleware.go    # Middleware JWT
├── models/
│   ├── user.go          # Modelo de usuario
│   ├── task.go          # Modelo de tarea
│   ├── task_request.go  # DTOs de peticiones
│   ├── constants.go     # Constantes de la app
│   └── migrate.go       # Migraciones
├── routes/
│   └── routes.go        # Definición de rutas
├── tests/
│   ├── helpers_test.go  # Funciones auxiliares
│   ├── auth_test.go     # Tests de autenticación
│   ├── task_test.go     # Tests de tareas
│   └── validation_test.go   # Tests de validaciones
├── .env                 # Variables de entorno (no en git)
├── .env.example         # Ejemplo de variables
├── .gitignore           # Archivos ignorados por git
├── go.mod               # Dependencias del proyecto
├── go.sum               # Checksums de dependencias
├── main.go              # Punto de entrada
└── README.md            # Este archivo
```

---

## 📦 Dependencias (go.mod)

Las dependencias principales están definidas en `go.mod`:

```go
module go-task-manager-mvc

go 1.21

require (
    github.com/gin-gonic/gin v1.10.0
    github.com/golang-jwt/jwt/v5 v5.2.1
    github.com/joho/godotenv v1.5.1
    golang.org/x/crypto v0.23.0
    gorm.io/driver/mysql v1.5.7
    gorm.io/gorm v1.25.12
)

require (
    github.com/stretchr/testify v1.9.0 // testing
)
```

### Instalar todas las dependencias

```bash
go mod download
```

---

## 🔧 Solución de Problemas

### Error: "Error al conectar con la base de datos"

**Causa:** Credenciales incorrectas o MySQL no está corriendo

**Solución:**
```bash
# Verificar que MySQL está corriendo
sudo systemctl status mysql  # Linux
# o
mysql.server status          # macOS

# Verificar credenciales en .env
cat .env
```

### Error: "JWT_SECRET no configurado"

**Causa:** Falta la variable JWT_SECRET en `.env`

**Solución:**
```bash
echo "JWT_SECRET=tu_clave_secreta_super_segura" >> .env
```

### Error: "Duplicate entry for key 'users.uni_users_email'"

**Causa:** El email ya está registrado

**Solución:** Usa otro email o elimina el usuario existente

### Tests fallan

**Causa:** Base de datos no está configurada o hay datos residuales

**Solución:**
```bash
# Limpiar base de datos de test
mysql -u root -p tasks_db < scripts/clean_test_data.sql
```

---

## 🤝 Contribuir

¡Las contribuciones son bienvenidas! Si deseas contribuir:

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

### Guías de contribución

- Sigue las convenciones de código de Go
- Asegúrate de que todos los tests pasen
- Agrega tests para nuevas funcionalidades
- Actualiza la documentación si es necesario

---

## 🎯 Roadmap

- [x] Autenticación JWT
- [x] CRUD de tareas
- [x] Testing completo
- [x] Validaciones robustas
- [ ] Documentación con Swagger
- [ ] Dockerfile
- [ ] CI/CD con GitHub Actions
- [ ] Deploy en Render/Railway
- [ ] Filtros avanzados (por fecha, prioridad)
- [ ] Paginación de resultados
- [ ] Rate limiting
- [ ] Logs estructurados

---

## 👨‍💻 Autor

**Stormdead**
- GitHub: [@Stormdead](https://github.com/Stormdead)

---

## 🙏 Agradecimientos

- [Gin Framework](https://gin-gonic.com/) - Framework web
- [GORM](https://gorm.io/) - ORM increíble
- [JWT-Go](https://github.com/golang-jwt/jwt) - Manejo de tokens
- Comunidad de Go por el excelente ecosistema

---

## 📚 Recursos Adicionales

- [Documentación de Go](https://golang.org/doc/)
- [Tutorial de Gin](https://gin-gonic.com/docs/)
- [GORM Guides](https://gorm.io/docs/)
- [JWT Introduction](https://jwt.io/introduction)

---

**⭐ Si este proyecto te fue útil, considera darle una estrella en GitHub!**

---