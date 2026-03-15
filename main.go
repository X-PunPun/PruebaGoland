package main

import (
	"gamevault/config"
	"gamevault/database"
	"gamevault/handlers"
	"gamevault/repositories"
	"gamevault/server"
	"gamevault/services"
)

func main() {
	// 1. Cargar configuración
	cfg := config.Load()

	// 2. Conectar a Base de Datos
	db := database.Connect(cfg.DBConnString)
	defer db.Close()

	// 3. Crear instancias (Inyección de dependencias manual)
	gameRepo := repositories.NewGameRepository(db)
	gameService := services.NewGameService(gameRepo, cfg)
	gameHandler := handlers.NewGameHandler(gameService)

	// 4. Configurar Servidor y Rutas
	app := server.NewApp()

	// Endpoints RAWG (Proxy)
	app.Get("/api/search", gameHandler.SearchRAWG)
	app.Get("/api/games/{id}", gameHandler.GetRAWGGame)

	// Endpoints Locales (CRUD)
	app.Get("/api/library", gameHandler.GetLibrary)
	app.Post("/api/library", gameHandler.AddToLibrary)
	app.Put("/api/library/{id}", gameHandler.UpdateGame)
	app.Delete("/api/library/{id}", gameHandler.DeleteGame)
	app.Get("/api/library/stats", gameHandler.GetStats)

	// 5. Levantar Servidor
	server.Run(app, cfg)
}
