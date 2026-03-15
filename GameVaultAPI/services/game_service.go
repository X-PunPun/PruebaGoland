package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"gamevault/config"
	"gamevault/models"
	"gamevault/repositories"
	"net/http"
)

type GameService interface {
	SearchRAWG(query string) (*models.RawgSearchResponse, error)
	GetRAWGGame(id string) (*models.RawgGame, error)
	AddToLibrary(game models.GameLibrary) error
	GetLibrary(status string) ([]models.GameLibrary, error)
	UpdateGame(id int, data models.GameUpdateDTO) error
	DeleteGame(id int) error
	GetStats() (*models.LibraryStats, error)
}

type gameService struct {
	repo   repositories.GameRepository
	config *config.Config
}

func NewGameService(repo repositories.GameRepository, cfg *config.Config) GameService {
	return &gameService{repo: repo, config: cfg}
}

func (s *gameService) SearchRAWG(query string) (*models.RawgSearchResponse, error) {
	url := fmt.Sprintf("%s/games?key=%s&search=%s", s.config.RawgBaseURL, s.config.RawgAPIKey, query)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("rawg_api_error")
	}
	defer resp.Body.Close()

	var result models.RawgSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *gameService) GetRAWGGame(id string) (*models.RawgGame, error) {
	url := fmt.Sprintf("%s/games/%s?key=%s", s.config.RawgBaseURL, id, s.config.RawgAPIKey)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		if resp != nil && resp.StatusCode == 404 {
			return nil, errors.New("not_found")
		}
		return nil, errors.New("rawg_api_error")
	}
	defer resp.Body.Close()

	var result models.RawgGame
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *gameService) AddToLibrary(game models.GameLibrary) error {
	// Validaciones
	if game.RawgID == 0 || game.Title == "" {
		return errors.New("bad_request")
	}

	exists, err := s.repo.CheckExistsByRawgID(game.RawgID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("conflict")
	}

	return s.repo.Create(&game)
}

func (s *gameService) GetLibrary(status string) ([]models.GameLibrary, error) {
	return s.repo.GetAll(status)
}

func (s *gameService) UpdateGame(id int, data models.GameUpdateDTO) error {
	// Verificar existencia
	game, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if game == nil {
		return errors.New("not_found")
	}

	// Validar score si se envía
	if data.PersonalScore != nil && !ValidatePersonalScore(*data.PersonalScore) {
		return errors.New("bad_request")
	}

	// Validar status si se envía
	if data.Status != nil && !ValidateStatus(*data.Status) {
		return errors.New("bad_request")
	}

	// // Validar score si se envía
	// if data.PersonalScore != nil && (*data.PersonalScore < 1 || *data.PersonalScore > 10) {
	// 	return errors.New("bad_request")
	// }

	// // Validar status si se envía
	// if data.Status != nil {
	// 	st := *data.Status
	// 	if st != "completado" && st != "jugando" && st != "pendiente" && st != "abandonado" {
	// 		return errors.New("bad_request")
	// 	}
	// }

	return s.repo.Update(id, data)
}

func (s *gameService) DeleteGame(id int) error {
	return s.repo.Delete(id)
}

func (s *gameService) GetStats() (*models.LibraryStats, error) {
	return s.repo.GetStats()
}

// --- Funciones auxiliares para validación (Extraídas para Testing) ---

// ValidateStatus verifica que el estado sea uno de los permitidos
func ValidateStatus(status string) bool {
	validStatuses := []string{"pendiente", "jugando", "completado", "abandonado"}
	for _, s := range validStatuses {
		if status == s {
			return true
		}
	}
	return false
}

// ValidatePersonalScore verifica que el puntaje esté entre 1 y 10
func ValidatePersonalScore(score int) bool {
	return score >= 1 && score <= 10
}
