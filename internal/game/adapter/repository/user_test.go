package repository

import (
	"errors"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql/model"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/repository/testdata"
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
	"testing"
)

func TestNew(t *testing.T) {
	connection := testdata.Connection{}
	repository := User{&connection}
	user := domain.User{
		Id:   "id",
		Name: "name",
	}
	err := repository.New(user)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}

	db, err := connection.Open()

	modelUser := model.User{Id: user.Id}
	db.Model(modelUser).Find(&modelUser)
	if modelUser.Name != user.Name {
		t.Error("failed to assert that user was added")
	}

	db.Delete(modelUser)
}

func TestNewErr(t *testing.T) {
	connection := testdata.Connection{}
	repository := User{&connection}
	testdata.ConnectionErr = errors.New("error")
	defer testdata.Reset()

	user := domain.User{
		Id:   "id",
		Name: "name",
	}
	err := repository.New(user)
	if err != testdata.ConnectionErr {
		t.Error("failed to assert that an error occurred")
	}
}

func TestGetAll(t *testing.T) {
	connection := testdata.Connection{}
	repository := User{&connection}
	users, err := repository.GetAll()
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if len(users) != 0 {
		t.Error("failed to assert that no users exists")
	}

	db, _ := connection.Open()

	user1 := model.User{
		Id:   "id-1",
		Name: "name",
	}
	db.Save(user1)
	defer db.Delete(user1)
	user2 := model.User{
		Id:   "id-2",
		Name: "name",
	}
	db.Save(user2)
	defer db.Delete(user2)

	users, err = repository.GetAll()
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if len(users) != 2 {
		t.Error("failed to assert number of users")
	}
	for _, u := range users {
		if len(u.Id) == 0 {
			t.Error("failed to assert that ID was set")
		}
		if len(u.Name) == 0 {
			t.Error("failed to assert that ID was set")
		}
		if u.Friends != nil {
			t.Error("failed to assert that no friends were fetched")
		}
	}
}

func TestGetAllErr(t *testing.T) {
	connection := testdata.Connection{}
	repository := User{&connection}
	testdata.ConnectionErr = errors.New("error")
	defer testdata.Reset()
	_, err := repository.GetAll()
	if err != testdata.ConnectionErr {
		t.Error("failed to assert that an error occurred")
	}
}

func TestGetByIds(t *testing.T) {
	connection := testdata.Connection{}
	repository := User{&connection}
	users, err := repository.GetByIds([]string{})
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if len(users) != 0 {
		t.Error("failed to assert that no users exists")
	}

	db, _ := connection.Open()

	user1 := model.User{
		Id:   "id-1",
		Name: "name",
	}
	db.Save(user1)
	defer db.Delete(user1)
	user2 := model.User{
		Id:   "id-2",
		Name: "name",
	}
	db.Save(user2)
	defer db.Delete(user2)

	users, err = repository.GetByIds([]string{user1.Id})
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if len(users) != 1 {
		t.Error("failed to assert number of users")
	} else {
		user := users[0]
		if user.Id != user1.Id {
			t.Error("failed to assert user1 ID")
		}
		if user.Name != user1.Name {
			t.Error("failed to assert user1 name")
		}
	}
	users, err = repository.GetByIds([]string{user1.Id, user2.Id})

	for _, u := range users {
		if len(u.Id) == 0 {
			t.Error("failed to assert that ID was set")
		}
		if len(u.Name) == 0 {
			t.Error("failed to assert that ID was set")
		}
		if u.Friends != nil {
			t.Error("failed to assert that no friends were fetched")
		}
	}
}

func TestGetByIdsErr(t *testing.T) {
	connection := testdata.Connection{}
	repository := User{&connection}
	testdata.ConnectionErr = errors.New("error")
	defer testdata.Reset()
	_, err := repository.GetByIds([]string{"id"})
	if err != testdata.ConnectionErr {
		t.Error("failed to assert that an error occurred")
	}
}
