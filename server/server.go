package server

import (
	"gamevault/config"
	"log"
	"net/http"
)

func Run(app *App, cfg *config.Config) {
	log.Printf("Iniciando servidor en puerto :%s", cfg.Port)

	err := http.ListenAndServe(":"+cfg.Port, app.GetMux())
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
