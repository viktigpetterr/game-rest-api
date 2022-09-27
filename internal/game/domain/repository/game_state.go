package repository

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
)

type IGameState interface {
	New(gameState domain.GameState) error
	Update(gameState domain.GameState) (domain.GameState, error)
	GetByUserId(id string) (domain.GameState, error)
	GetScoreByUserId(id string) (int, error)
}
