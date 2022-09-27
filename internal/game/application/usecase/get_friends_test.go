package usecase

import (
	"errors"
	"github.com/viktigpetterr/game-rest-api/internal/game/application/usecase/testdata/repository"
	"testing"
)

func TestGetFriends(t *testing.T) {
	mock := repository.FriendMock{}
	args := GetFriendsArgs{UserId: "id"}
	friends, err := GetFriends(mock, args)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if len(friends) == 0 {
		t.Error("failed to asser that friends were fetched")
	}
	friend := friends[0]
	if friend.Name != "name" {
		t.Error("failed to assert that user name was set correctly")
	}
	if friend.Id != "id" {
		t.Error("failed to assert that user ID was set correctly")
	}
	if friend.HighScore != 0 {
		t.Error("failed to assert that game state score was set correctly")
	}
}

func TestGetFriendsWithErr(t *testing.T) {
	mock := repository.FriendMock{}
	args := GetFriendsArgs{UserId: "id"}
	repository.FriendErr = errors.New("test")
	_, err := GetFriends(mock, args)
	if err != repository.FriendErr {
		t.Error("failed to assert that an error occurred")
	}
	repository.FriendErr = nil
}
