package server

import "net/http"

type HandlerFunc func(c *Context)

func (a *App) registerRoute(method, path string, handler HandlerFunc) {
	// Patrón de ruteo de Go 1.22+ (ej: "GET /api/users")
	pattern := method + " " + path
	a.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{W: w, R: r}
		handler(ctx)
	})
}

// Helpers para hacer la configuración de rutas más limpia
func (a *App) Get(path string, handler HandlerFunc) {
	a.registerRoute(http.MethodGet, path, handler)
}
func (a *App) Post(path string, handler HandlerFunc) {
	a.registerRoute(http.MethodPost, path, handler)
}
func (a *App) Put(path string, handler HandlerFunc) {
	a.registerRoute(http.MethodPut, path, handler)
}
func (a *App) Delete(path string, handler HandlerFunc) {
	a.registerRoute(http.MethodDelete, path, handler)
}
