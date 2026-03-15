package models

// GameLibrary representa la entidad en la base de datos local
type GameLibrary struct {
	ID            int    `json:"id"`
	RawgID        int    `json:"rawg_id"`
	Title         string `json:"title"`
	Genre         string `json:"genre"`
	Platform      string `json:"platform"`
	CoverURL      string `json:"cover_url"`
	PersonalNote  string `json:"personal_note"`
	PersonalScore int    `json:"personal_score"`
	Status        string `json:"status"`
	AddedAt       string `json:"added_at"`
}

// GameUpdateDTO se usa para el PUT (parcial)
type GameUpdateDTO struct {
	PersonalNote  *string `json:"personal_note"`
	PersonalScore *int    `json:"personal_score"`
	Status        *string `json:"status"`
}

// LibraryStats representa las estadísticas
type LibraryStats struct {
	Total        int            `json:"total"`
	ByStatus     map[string]int `json:"by_status"`
	AverageScore float64        `json:"average_score"`
}
