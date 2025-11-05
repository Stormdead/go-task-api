<<<<<<< HEAD
# go-task-api
=======
🦫 Proyecto 1: API REST de Gestión de Tareas – Go puro (backend sólido)
🎯 Objetivo:

Desarrollar una API RESTful completa usando Go (Golang) que permita gestionar tareas (crear, leer, actualizar, eliminar) con autenticación de usuarios.

🧩 Tecnologías:

Lenguaje: Go

Framework: Gin o Fiber

Base de datos: PostgreSQL

Autenticación: JWT

ORM: GORM o sqlx

Testing: Pruebas unitarias básicas con testing y httptest

🏗️ Estructura recomendada:
/cmd
  /api
    main.go
/internal
  /controllers
  /models
  /repositories
  /services
  /middlewares
/config
/database.go

🔑 Funcionalidades:

Registro e inicio de sesión de usuarios.

CRUD completo de tareas (título, descripción, estado, fecha límite).

Filtros de búsqueda (por usuario, por estado o por fecha).

Protección con JWT (solo el usuario autenticado ve sus tareas).

Logs de eventos (crear, editar, eliminar).

🚀 Extras para destacar:

Documentación de la API con Swagger.

Dockerfile para levantar el backend fácilmente.

Deploy gratuito en Render o Railway.

💬 Cómo presentarlo:

API REST desarrollada en Go con arquitectura modular, autenticación JWT y persistencia en PostgreSQL. Diseñada para demostrar buenas prácticas en diseño de servicios backend escalables y mantenibles.
>>>>>>> 06f82c9 (Versión inicial del Task Manager MVC con JWT y MySQL)
