package server

import "net/http"

type App struct {
	mux *http.ServeMux
}

func NewApp() *App {
	return &App{
		mux: http.NewServeMux(),
	}
}

// Exponemos el mux subyacente para que el servidor lo use
func (a *App) GetMux() *http.ServeMux {
	return a.mux
}
