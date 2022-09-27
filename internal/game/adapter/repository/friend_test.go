package repository

import (
	"errors"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql/model"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/repository/testdata"
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
	"testing"
)

func TestPutByUserId(t *testing.T) {
	connection := testdata.Connection{}
	repository := Friend{&connection}
	db, _ := connection.Open()

	userModel := model.User{
		Id:   "id",
		Name: "name",
	}
	userState := model.GameState{
		GamesPlayed: 0,
		Score:       0,
	}
	db.Save(userModel)
	userState.UserId = userModel.Id
	db.Save(userState)

	friendModel := model.User{
		Id:   "friend-id",
		Name: "friend",
	}
	db.Save(friendModel)
	userState.UserId = friendModel.Id
	db.Save(userState)

	defer db.Delete([]model.User{userModel, friendModel})

	friend := domain.User{Id: friendModel.Id}
	err := repository.PutByUserId([]domain.User{friend}, userModel.Id)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	friends, _ := repository.GetByUserId(userModel.Id)
	if len(friends) != 1 {
		t.Error("failed to assert number of friends")
	} else {
		if friends[0].Id != friend.Id {
			t.Error("failed to assert friend ID")
		}
	}
}

func TestPutByUserIdErr(t *testing.T) {
	connection := testdata.Connection{}
	repository := Friend{&connection}
	testdata.ConnectionErr = errors.New("error")
	defer testdata.Reset()
	err := repository.PutByUserId(nil, "")
	if err != testdata.ConnectionErr {
		t.Error("failed to assert that an error occurred")
	}
}

func TestGetByUserIdNone(t *testing.T) {
	connection := testdata.Connection{}
	repository := Friend{&connection}
	friends, err := repository.GetByUserId("0")
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if len(friends) != 0 {
		t.Error("failed to assert that no friends were found")
	}
}

func TestGetByUserIdErr(t *testing.T) {
	connection := testdata.Connection{}
	repository := Friend{&connection}
	testdata.ConnectionErr = errors.New("error")
	defer testdata.Reset()
	_, err := repository.GetByUserId("")
	if err != testdata.ConnectionErr {
		t.Error("failed to assert that an error occurred")
	}
}

func TestGetByUserIdId(t *testing.T) {
	connection := testdata.Connection{}
	repository := Friend{&connection}
	db, err := connection.Open()

	user := model.User{
		Id:   "id",
		Name: "name",
	}
	userState := model.GameState{
		GamesPlayed: 0,
		Score:       0,
	}
	db.Save(user)
	userState.UserId = user.Id
	db.Save(userState)

	friend := model.User{
		Id:   "friend-id",
		Name: "friend",
	}
	db.Save(friend)
	userState.UserId = friend.Id
	db.Save(userState)

	user.Friends = []*model.User{&friend}
	db.Save(user)

	defer db.Delete([]model.User{user, friend})

	friends, err := repository.GetByUserId("id")
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if len(friends) != 1 {
		t.Error("failed to assert the number of friends found")
	} else {
		f := friends[0]
		if f.Id != "friend-id" {
			t.Error("failed to assert friend")
		}
		if f.Name != "friend" {
			t.Error("failed to assert friend name")
		}
		if f.HighScore != 0 {
			t.Error("failed to asser friend's high score")
		}
	}
}
