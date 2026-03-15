package server

import "net/http"

// Context encapsula Request y ResponseWriter
type Context struct {
	W http.ResponseWriter
	R *http.Request
}

// Param obtiene una variable de ruta (Go 1.22+)
func (c *Context) Param(key string) string {
	return c.R.PathValue(key)
}

// Query obtiene un parámetro de la URL (?status=...)
func (c *Context) Query(key string) string {
	return c.R.URL.Query().Get(key)
}
