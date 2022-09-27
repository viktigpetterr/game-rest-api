package repository

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
)

type GameStateMock struct{}

var GameStateGetErr error
var GameStateUpdateErr error

var GameStateScore int

func (_ GameStateMock) New(_ domain.GameState) error {
	return GameStateGetErr
}

func (_ GameStateMock) Update(gameState domain.GameState) (domain.GameState, error) {
	return gameState, GameStateUpdateErr
}

func (_ GameStateMock) GetByUserId(id string) (domain.GameState, error) {
	return domain.GameState{
		UserId:      id,
		GamesPlayed: 0,
		Score:       0,
	}, GameStateGetErr
}

func (_ GameStateMock) GetScoreByUserId(_ string) (int, error) {
	return GameStateScore, GameStateGetErr
}

func (_ GameStateMock) Reset() {
	GameStateGetErr = nil
	GameStateUpdateErr = nil
}
