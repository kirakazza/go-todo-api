package main

import (
	"log"
	"os"

	"github.com/kirakazza/go-todo-api/internal/handler"
	"github.com/kirakazza/go-todo-api/internal/middleware"
	"github.com/kirakazza/go-todo-api/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env")
	}

	repository.InitDB() // инициализация базы

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.POST("/register", handler.RegisterHandler)
	router.POST("/login", handler.LoginHandler)

	auth := middleware.AuthMiddleware()

	authorized := router.Group("/")
	authorized.Use(auth)
	{
		authorized.GET("/todos", handler.GetTodosHandler)
		authorized.POST("/todos", handler.CreateTodoHandler)
		authorized.PUT("/todos/:id", handler.UpdateTodoHandler)
		authorized.DELETE("/todos/:id", handler.DeleteTodoHandler)
		// и другие приватные роуты
	}

	log.Println("Сервер запущен на порту", port)
	router.Run(":" + port)
}
