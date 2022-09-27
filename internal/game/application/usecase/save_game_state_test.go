package usecase

import (
	"errors"
	"github.com/viktigpetterr/game-rest-api/internal/game/application/usecase/testdata/repository"
	"testing"
)

func TestSaveGameState(t *testing.T) {
	mock := repository.GameStateMock{}
	args := SaveGameArgs{
		UserId:      "id",
		GamesPlayed: 1,
		Score:       50,
	}
	state, err := SaveGameState(mock, args)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if state.UserId != args.UserId {
		t.Error("failed to assert state user ID")
	}
	if state.GamesPlayed != args.GamesPlayed {
		t.Error("failed to assert played games")
	}
	if state.Score != args.Score {
		t.Error("failed to assert state score")
	}
}

func TestSaveGameStateUnalteredScore(t *testing.T) {
	mock := repository.GameStateMock{}
	repository.GameStateScore = 50
	args := SaveGameArgs{
		UserId:      "id",
		GamesPlayed: 1,
		Score:       0,
	}
	state, _ := SaveGameState(mock, args)
	if state.UserId != args.UserId {
		t.Error("failed to assert state user ID")
	}
	if state.GamesPlayed != args.GamesPlayed {
		t.Error("failed to assert played games")
	}
	if state.Score != repository.GameStateScore {
		t.Error("failed to assert state score")
	}
}

func TestSaveGameStateWithGetErr(t *testing.T) {
	mock := repository.GameStateMock{}
	repository.GameStateGetErr = errors.New("test")
	defer mock.Reset()
	args := SaveGameArgs{
		UserId:      "id",
		GamesPlayed: 1,
		Score:       0,
	}
	_, err := SaveGameState(mock, args)
	if err != repository.GameStateGetErr {
		t.Error("failed to assert that an error occurred")
	}
}

func TestSaveGameStateWithUpdateErr(t *testing.T) {
	mock := repository.GameStateMock{}
	repository.GameStateGetErr = nil
	repository.GameStateUpdateErr = errors.New("test")
	defer mock.Reset()
	args := SaveGameArgs{
		UserId:      "id",
		GamesPlayed: 1,
		Score:       0,
	}
	_, err := SaveGameState(mock, args)
	if err != repository.GameStateUpdateErr {
		t.Error("failed to assert that an error occurred")
	}
}
