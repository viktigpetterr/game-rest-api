package usecase

import (
	"errors"
	"github.com/viktigpetterr/game-rest-api/internal/game/application/usecase/testdata/repository"
	"testing"
)

func TestLoadGameState(t *testing.T) {
	mock := repository.GameStateMock{}
	args := LoadGameArgs{UserId: "id"}
	gameState, err := LoadGameState(mock, args)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if gameState.UserId != args.UserId {
		t.Error("failed to assert user ID")
	}
	if gameState.GamesPlayed != 0 {
		t.Error("failed to assert number of played games")
	}
	if gameState.Score != 0 {
		t.Error("failed to assert score")
	}
}

func TestLoadGameStateWithErr(t *testing.T) {
	mock := repository.GameStateMock{}
	repository.GameStateGetErr = errors.New("test")
	defer mock.Reset()
	args := LoadGameArgs{UserId: "id"}
	_, err := LoadGameState(mock, args)
	if err != repository.GameStateGetErr {
		t.Error("failed to assert that an error occurred")
	}
}
