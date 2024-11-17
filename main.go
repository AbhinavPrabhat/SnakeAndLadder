package main

import (
    "task/controllers"
    "task/service"
	"task/docs" 
    "github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
    echoSwagger "github.com/swaggo/echo-swagger"
)
// @title Snake and Ladder Game API
// @version 1.0
// @description A RESTful API for playing Snake and Ladder game
// @host localhost:8085
// @BasePath /
func main() {
    // Create new Echo instance
    e := echo.New()

    // Middleware
    e.Use(echomiddleware.Logger())
    e.Use(echomiddleware.Recover())
    e.Use(echomiddleware.CORS())

	// Initialize Swagger metadata
    docs.SwaggerInfo.Title = "Snake and Ladder Game API"
    docs.SwaggerInfo.Description = "A RESTful API for playing Snake and Ladder game"
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost:8085"
    docs.SwaggerInfo.BasePath = "/"

    // Initialize services and controllers
    gameService := service.NewGameService()
    gameController := controllers.NewGameController(gameService)
	 // Swagger endpoint
    e.GET("/swagger/*", echoSwagger.WrapHandler)

    // Routes
    // Lobby routes
    e.POST("/lobby/create", gameController.CreateLobby)
    e.POST("/lobby/join", gameController.JoinLobby)

    // Game routes
    e.POST("/game/roll", gameController.RollDice)
    e.GET("/game/:gameId", gameController.GetGame)

    // Start server
    e.Logger.Fatal(e.Start(":8085"))
}