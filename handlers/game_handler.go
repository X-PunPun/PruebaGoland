package handlers

import (
	"encoding/json"
	"gamevault/models"
	"gamevault/server"
	"gamevault/services"
	"log"
	"strconv"
)

type GameHandler struct {
	service services.GameService
}

func NewGameHandler(s services.GameService) *GameHandler {
	return &GameHandler{service: s}
}

// GET /api/search?q={nombre}
func (h *GameHandler) SearchRAWG(c *server.Context) {
	q := c.Query("q")
	if q == "" {
		RespondError(c.W, 400, "400", "Parámetro 'q' es requerido")
		return
	}

	result, err := h.service.SearchRAWG(q)
	if err != nil {
		log.Println("Error RAWG:", err)
		RespondError(c.W, 502, "502", "Error comunicándose con RAWG API")
		return
	}

	RespondJSON(c.W, 200, result.Results)
}

// GET /api/games/{id}
func (h *GameHandler) GetRAWGGame(c *server.Context) {
	id := c.Param("id")
	result, err := h.service.GetRAWGGame(id)
	if err != nil {
		if err.Error() == "not_found" {
			RespondError(c.W, 404, "404", "Juego no encontrado en RAWG")
			return
		}
		log.Println("Error RAWG:", err)
		RespondError(c.W, 502, "502", "Error comunicándose con RAWG API")
		return
	}

	RespondJSON(c.W, 200, result)
}

// GET /api/library
func (h *GameHandler) GetLibrary(c *server.Context) {
	status := c.Query("status")
	games, err := h.service.GetLibrary(status)
	if err != nil {
		log.Println("Error BD:", err)
		RespondError(c.W, 500, "500", "Error interno del servidor")
		return
	}

	RespondJSON(c.W, 200, games)
}

// POST /api/library
func (h *GameHandler) AddToLibrary(c *server.Context) {
	var game models.GameLibrary
	if err := json.NewDecoder(c.R.Body).Decode(&game); err != nil {
		RespondError(c.W, 400, "400", "Cuerpo de petición inválido")
		return
	}

	err := h.service.AddToLibrary(game)
	if err != nil {
		if err.Error() == "bad_request" {
			RespondError(c.W, 400, "400", "Faltan campos obligatorios (rawg_id, title)")
			return
		}
		if err.Error() == "conflict" {
			RespondError(c.W, 409, "409", "El juego ya existe en la colección")
			return
		}
		log.Println("Error BD:", err)
		RespondError(c.W, 500, "500", "Error interno del servidor")
		return
	}

	RespondJSON(c.W, 201, map[string]string{"message": "Juego agregado correctamente"})
}

// PUT /api/library/{id}
func (h *GameHandler) UpdateGame(c *server.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondError(c.W, 400, "400", "ID inválido")
		return
	}

	var data models.GameUpdateDTO
	if err := json.NewDecoder(c.R.Body).Decode(&data); err != nil {
		RespondError(c.W, 400, "400", "Cuerpo de petición inválido")
		return
	}

	err = h.service.UpdateGame(id, data)
	if err != nil {
		if err.Error() == "not_found" {
			RespondError(c.W, 404, "404", "Juego no encontrado")
			return
		}
		if err.Error() == "bad_request" {
			RespondError(c.W, 400, "400", "Datos inválidos (score fuera de rango o status incorrecto)")
			return
		}
		log.Println("Error BD:", err)
		RespondError(c.W, 500, "500", "Error interno del servidor")
		return
	}

	RespondJSON(c.W, 200, map[string]string{"message": "Juego actualizado"})
}

// DELETE /api/library/{id}
func (h *GameHandler) DeleteGame(c *server.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondError(c.W, 400, "400", "ID inválido")
		return
	}

	err = h.service.DeleteGame(id)
	if err != nil {
		if err.Error() == "not_found" {
			RespondError(c.W, 404, "404", "Juego no encontrado")
			return
		}
		log.Println("Error BD:", err)
		RespondError(c.W, 500, "500", "Error interno del servidor")
		return
	}

	RespondJSON(c.W, 204, nil)
}

// GET /api/library/stats
func (h *GameHandler) GetStats(c *server.Context) {
	stats, err := h.service.GetStats()
	if err != nil {
		log.Println("Error BD:", err)
		RespondError(c.W, 500, "500", "Error interno del servidor")
		return
	}

	RespondJSON(c.W, 200, stats)
}
