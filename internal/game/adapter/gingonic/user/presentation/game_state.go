package presentation

import "github.com/viktigpetterr/game-rest-api/internal/game/domain"

type JsonGameState struct {
	GamesPlayed int `json:"gamesPlayed"`
	Score       int `json:"score"`
}

func MakeJsonGameState(gameState domain.GameState) JsonGameState {
	return JsonGameState{
		gameState.GamesPlayed,
		gameState.Score,
	}
}
