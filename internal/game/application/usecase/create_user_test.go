package usecase

import (
	"errors"
	"github.com/viktigpetterr/game-rest-api/internal/game/application/usecase/testdata/repository"
	"github.com/viktigpetterr/game-rest-api/internal/game/application/usecase/testdata/service"
	"testing"
)

func TestCreateUser(t *testing.T) {
	uuidMock := service.UuidMock{}
	gameStateRepoMock := repository.GameStateMock{}
	userRepoMock := repository.UserMock{}

	args := CreateUserArgs{Name: "test"}

	user, err := CreateUser(
		gameStateRepoMock,
		userRepoMock,
		uuidMock,
		args,
	)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if user.Name != args.Name {
		t.Error("failed to assert user name")
	}
	if user.Id != uuidMock.New() {
		t.Error("failed to assert user ID")
	}
	gameState := user.GameState
	if gameState.Score != 0 {
		t.Error("failed to assert that user had game state with score 0")
	}
	if gameState.GamesPlayed != 0 {
		t.Error("failed to assert that user had game state without played games")
	}
}

func TestCreateUserUserErr(t *testing.T) {
	uuidMock := service.UuidMock{}
	userRepoMock := repository.UserMock{}
	gameStateRepoMock := repository.GameStateMock{}

	repository.GameStateGetErr = errors.New("test")
	defer gameStateRepoMock.Reset()
	args := CreateUserArgs{Name: "test"}
	_, err := CreateUser(
		gameStateRepoMock,
		userRepoMock,
		uuidMock,
		args,
	)
	if err != repository.GameStateGetErr {
		t.Error("failed to assert an error occurred")
	}
}

func TestCreateUserGameStateErr(t *testing.T) {
	uuidMock := service.UuidMock{}
	userRepoMock := repository.UserMock{}
	repository.UserErr = errors.New("test")

	args := CreateUserArgs{Name: "test"}

	_, err := CreateUser(
		nil,
		userRepoMock,
		uuidMock,
		args,
	)
	if err != repository.UserErr {
		t.Error("failed to assert an error occurred")
	}

	repository.UserErr = nil
}
