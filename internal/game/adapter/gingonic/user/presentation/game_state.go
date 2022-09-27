package presentation

import "github.com/viktigpetterr/game-rest-api/internal/game/domain"

type GameState struct {
	GamesPlayed int `json:"gamesPlayed"`
	Score       int `json:"score"`
}

func MakeGameState(gameState domain.GameState) GameState {
	return GameState{
		gameState.GamesPlayed,
		gameState.Score,
	}
}
