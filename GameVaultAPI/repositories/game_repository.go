package repositories

import (
	"database/sql"
	"errors"
	"gamevault/models"
)

type GameRepository interface {
	Create(game *models.GameLibrary) error
	GetAll(statusFilter string) ([]models.GameLibrary, error)
	GetByID(id int) (*models.GameLibrary, error)
	Update(id int, data models.GameUpdateDTO) error
	Delete(id int) error
	GetStats() (*models.LibraryStats, error)
	CheckExistsByRawgID(rawgID int) (bool, error)
}

type gameRepository struct {
	db *sql.DB
}

func NewGameRepository(db *sql.DB) GameRepository {
	return &gameRepository{db: db}
}

func (r *gameRepository) Create(g *models.GameLibrary) error {
	query := `
		INSERT INTO game_library (rawg_id, title, genre, platform, cover_url, status) 
		VALUES (@p1, @p2, @p3, @p4, @p5, @p6)`

	// Por defecto pendiente
	status := g.Status
	if status == "" {
		status = "pendiente"
	}

	_, err := r.db.Exec(query, g.RawgID, g.Title, g.Genre, g.Platform, g.CoverURL, status)
	return err
}

func (r *gameRepository) CheckExistsByRawgID(rawgID int) (bool, error) {
	var exists bool
	query := "SELECT CASE WHEN EXISTS (SELECT 1 FROM game_library WHERE rawg_id = @p1) THEN 1 ELSE 0 END"
	err := r.db.QueryRow(query, rawgID).Scan(&exists)
	return exists, err
}

func (r *gameRepository) GetAll(statusFilter string) ([]models.GameLibrary, error) {
	query := "SELECT id, rawg_id, title, genre, platform, cover_url, ISNULL(personal_note, ''), ISNULL(personal_score, 0), status FROM game_library"
	var rows *sql.Rows
	var err error

	if statusFilter != "" {
		query += " WHERE status = @p1"
		rows, err = r.db.Query(query, statusFilter)
	} else {
		rows, err = r.db.Query(query)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.GameLibrary
	for rows.Next() {
		var g models.GameLibrary
		err := rows.Scan(&g.ID, &g.RawgID, &g.Title, &g.Genre, &g.Platform, &g.CoverURL, &g.PersonalNote, &g.PersonalScore, &g.Status)
		if err != nil {
			return nil, err
		}
		games = append(games, g)
	}
	if games == nil {
		games = []models.GameLibrary{} // Retornar array vacío en vez de null
	}
	return games, nil
}

func (r *gameRepository) GetByID(id int) (*models.GameLibrary, error) {
	query := "SELECT id, rawg_id, title, genre, platform, cover_url, ISNULL(personal_note, ''), ISNULL(personal_score, 0), status FROM game_library WHERE id = @p1"
	var g models.GameLibrary
	err := r.db.QueryRow(query, id).Scan(&g.ID, &g.RawgID, &g.Title, &g.Genre, &g.Platform, &g.CoverURL, &g.PersonalNote, &g.PersonalScore, &g.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No encontrado
		}
		return nil, err
	}
	return &g, nil
}

func (r *gameRepository) Update(id int, data models.GameUpdateDTO) error {
	// Actualización dinámica basada en qué campos llegaron (no nil)
	query := "UPDATE game_library SET "
	var params []interface{}
	paramCount := 1

	if data.PersonalNote != nil {
		query += "personal_note = @p" + string(rune('0'+paramCount)) + ", "
		params = append(params, *data.PersonalNote)
		paramCount++
	}
	if data.PersonalScore != nil {
		query += "personal_score = @p" + string(rune('0'+paramCount)) + ", "
		params = append(params, *data.PersonalScore)
		paramCount++
	}
	if data.Status != nil {
		query += "status = @p" + string(rune('0'+paramCount)) + ", "
		params = append(params, *data.Status)
		paramCount++
	}

	// Quitar la última coma y espacio
	query = query[:len(query)-2]

	// Agregar WHERE
	query += " WHERE id = @p" + string(rune('0'+paramCount))
	params = append(params, id)

	_, err := r.db.Exec(query, params...)
	return err
}

func (r *gameRepository) Delete(id int) error {
	query := "DELETE FROM game_library WHERE id = @p1"
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("not_found")
	}
	return err
}

func (r *gameRepository) GetStats() (*models.LibraryStats, error) {
	stats := &models.LibraryStats{
		ByStatus: make(map[string]int),
	}

	// Total y promedio
	queryAgg := "SELECT COUNT(*), ISNULL(AVG(CAST(personal_score AS FLOAT)), 0) FROM game_library WHERE personal_score IS NOT NULL AND personal_score > 0"
	err := r.db.QueryRow(queryAgg).Scan(&stats.Total, &stats.AverageScore)
	if err != nil {
		return nil, err
	}

	// Conteo por status
	queryStatus := "SELECT status, COUNT(*) FROM game_library GROUP BY status"
	rows, err := r.db.Query(queryStatus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err == nil {
			if status != "" {
				stats.ByStatus[status] = count
			}
		}
	}

	return stats, nil
}
