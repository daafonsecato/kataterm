# github.com/david8128/Kataterm-backend

Este proyecto es un backend desarrollado en Go que se conecta a una base de datos PostgreSQL. Contiene un CRM para manejar preguntas y un controlador para ejecutar tareas de validación a través de scripts de bash.

## Estructura del Proyecto

El proyecto tiene la siguiente estructura:

- `cmd/main.go`: Punto de entrada de la aplicación.
- `pkg/crm/question.go`: Maneja las operaciones de CRM relacionadas con las preguntas.
- `pkg/crm/question_test.go`: Pruebas unitarias para las funciones en `question.go`.
- `pkg/tasks/task.go`: Ejecuta y maneja tareas de validación.
- `pkg/tasks/task_test.go`: Pruebas unitarias para las funciones en `task.go`.
- `pkg/db/db.go`: Maneja la conexión y las operaciones de la base de datos PostgreSQL.
- `pkg/db/db_test.go`: Pruebas unitarias para las funciones en `db.go`.
- `scripts/validation.sh`: Script de bash que realiza tareas de validación.
- `go.mod` y `go.sum`: Manejo de dependencias de Go.

## Cómo Correr el Proyecto

Para correr el proyecto, necesitas tener Go y PostgreSQL instalados en tu máquina.

1. Clona el repositorio en tu máquina local.
2. Navega al directorio del proyecto.
3. Ejecuta `go run cmd/main.go` para iniciar la aplicación.

## Pruebas

Para correr las pruebas, navega al directorio del proyecto y ejecuta `go test ./...`.

## Contribuir

Si deseas contribuir a este proyecto, por favor haz un fork del repositorio, crea una nueva rama, y envía un pull request.