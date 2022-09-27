package usecase

import (
	"errors"
	"github.com/viktigpetterr/game-rest-api/internal/game/application/usecase/testdata/repository"
	"testing"
)

func TestUpdateFriends(t *testing.T) {
	userRepositoryMock := repository.UserMock{}
	friendRepositoryMock := repository.FriendMock{}
	args := UpdateFriendsArgs{
		UserId:  "id1",
		Friends: []string{"id2", "id3"},
	}
	err := UpdateFriends(userRepositoryMock, friendRepositoryMock, args)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
}

func TestUpdateFriendsWithUserErr(t *testing.T) {
	userRepositoryMock := repository.UserMock{}
	friendRepositoryMock := repository.FriendMock{}
	args := UpdateFriendsArgs{
		UserId:  "id1",
		Friends: []string{"id2", "id3"},
	}
	repository.UserErr = errors.New("test")
	err := UpdateFriends(userRepositoryMock, friendRepositoryMock, args)
	if err != repository.UserErr {
		t.Error("failed to assert that an error occurred")
	}
}

func TestUpdateFriendsWithFriendsErr(t *testing.T) {
	userRepositoryMock := repository.UserMock{}
	friendRepositoryMock := repository.FriendMock{}
	args := UpdateFriendsArgs{
		UserId:  "id1",
		Friends: []string{"id2", "id3"},
	}
	repository.UserErr = nil
	repository.FriendErr = errors.New("test")
	err := UpdateFriends(userRepositoryMock, friendRepositoryMock, args)
	if err != repository.FriendErr {
		t.Error("failed to assert that an error occurred")
	}
}
