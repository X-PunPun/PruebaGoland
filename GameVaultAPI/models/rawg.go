package models

// Estructuras para mapear la respuesta de RAWG
type RawgSearchResponse struct {
	Results []RawgGame `json:"results"`
}

type RawgGame struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Background string        `json:"background_image"`
	Rating     float64       `json:"rating"`
	Genres     []RawgGenre   `json:"genres"`
	Platforms  []RawgWrapper `json:"platforms"`
}

type RawgGenre struct {
	Name string `json:"name"`
}

type RawgWrapper struct {
	Platform RawgPlatform `json:"platform"`
}

type RawgPlatform struct {
	Name string `json:"name"`
}
