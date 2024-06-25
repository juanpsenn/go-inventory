package main

import (
	"log"

	"example.com/users/cmd/api/handlers/user"
	"example.com/users/internal/repositories/sqlite"
	userRepository "example.com/users/internal/repositories/sqlite/user"
	userService "example.com/users/internal/services/user"
	"github.com/gin-gonic/gin"
)

func main() {
	// Crear un router de Gin.
	router := gin.Default()

	dbSession, err := sqlite.Connect("./users.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbSession.Close()

	repository := userRepository.NewRepository(dbSession)

	userHandler := user.Handler{
		UserService: userService.NewService(
			userService.WithRepository(repository),
		),
	}

	// Ruta para registrar un nuevo usuario.
	router.POST("/users", userHandler.CreateUser)

	// Iniciar el servidor en el puerto 8080.
	router.Run(":8080")
}
