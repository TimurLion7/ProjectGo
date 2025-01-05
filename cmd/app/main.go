package main

import (
	"log"
	"myproject/internal/database"
	"myproject/internal/handlers"
	"myproject/internal/taskService"
	"myproject/internal/userService"
	"myproject/internal/web/tasks"
	"myproject/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	//database.DB.AutoMigrate(&taskService.Task{})

	TasksRepo := taskService.NewTaskRepository(database.DB)
	TasksService := taskService.NewTaskService(TasksRepo)
	TasksHandler := handlers.NewTaskHandler(TasksService)

	UsersRepo := userService.NewUserRepository(database.DB)
	UsersService := userService.NewUserService(UsersRepo)
	UsersHandler := handlers.NewUserHandler(UsersService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictTasksHandler := tasks.NewStrictHandler(TasksHandler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictTasksHandler)

	strictUsersHandler := users.NewStrictHandler(UsersHandler, nil) // тут будет ошибка
	users.RegisterHandlers(e, strictUsersHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}

}
