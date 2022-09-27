package presentation

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
	"testing"
)

func TestMakeGameState(t *testing.T) {
	domainGameState := domain.GameState{
		UserId:      "id",
		GamesPlayed: 1,
		Score:       2,
	}

	gameState := MakeGameState(domainGameState)
	if gameState.GamesPlayed != domainGameState.GamesPlayed {
		t.Error("failed to assert games played")
	}
	if gameState.Score != domainGameState.Score {
		t.Error("failed to assert score")
	}
}
