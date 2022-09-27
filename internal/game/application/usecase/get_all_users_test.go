package usecase

import (
	"errors"
	"github.com/viktigpetterr/game-rest-api/internal/game/application/usecase/testdata/repository"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	mock := repository.UserMock{}
	users, err := GetAllUsers(mock)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if len(users) == 0 {
		t.Error("failed to assert that all users were fetched")
	}
	user := users[0]
	if user.Name != "name" {
		t.Error("failed to assert that user name was set correctly")
	}
	if user.Id != "id" {
		t.Error("failed to assert that user ID was set correctly")
	}
	if user.GameState.UserId != "id" {
		t.Error("failed to assert that game state user ID was set correctly")
	}
	if user.GameState.Score != 0 {
		t.Error("failed to assert that game state score was set correctly")
	}
	if user.GameState.GamesPlayed != 0 {
		t.Error("failed to assert that game state games played was set correctly")
	}
}

func TestGetAllUsersWithErr(t *testing.T) {
	mock := repository.UserMock{}
	repository.UserErr = errors.New("test")
	_, err := GetAllUsers(mock)
	if err != repository.UserErr {
		t.Error("failed to assert that an error occurred")
	}
	repository.UserErr = nil
}
