package user

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/gingonic/user/presentation"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/service"
	"github.com/viktigpetterr/game-rest-api/internal/game/application/usecase/testdata/repository"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var recorder *httptest.ResponseRecorder
var c *gin.Context
var ctrl Controller

func setUp() {
	recorder = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(recorder)
	ctrl = Controller{
		userRepository:      repository.UserMock{},
		friendRepository:    repository.FriendMock{},
		gameStateRepository: repository.GameStateMock{},
		uuidService:         service.Uuid{},
	}
	repository.UserErr = nil
	repository.FriendErr = nil
	repository.GameStateGetErr = nil
	repository.GameStateUpdateErr = nil
}

func TestGetUsers(t *testing.T) {
	setUp()
	ctrl.GetUsers(c)
	body := getResponseBody(http.StatusOK, t)
	var users presentation.Users
	err := json.Unmarshal(body, &users)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if len(users.Users) != 1 {
		t.Error("failed to assert number of users")
	} else {
		if users.Users[0].Id != "id" {
			t.Error("failed to assert user ID")
		}
		if users.Users[0].Name != "name" {
			t.Error("failed to assert user name")
		}
	}
}

func TestGetUsersInternalError(t *testing.T) {
	setUp()
	repository.UserErr = errors.New("error")
	ctrl.GetUsers(c)
	assertEmptyResponse(http.StatusInternalServerError, t)
}

func TestPostUser(t *testing.T) {
	setUp()
	stringReader := strings.NewReader("{\"name\": \"name\"}")
	c.Request = &http.Request{Body: io.NopCloser(stringReader)}
	ctrl.PostUser(c)
	body := getResponseBody(http.StatusOK, t)
	user := presentation.User{}
	err := json.Unmarshal(body, &user)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if user.Name != "name" {
		t.Error("failed to assert name")
	}
}

func TestPostUserInternalError(t *testing.T) {
	setUp()
	stringReader := strings.NewReader("{\"name\": \"name\"}")
	c.Request = &http.Request{Body: io.NopCloser(stringReader)}
	repository.UserErr = errors.New("error")
	ctrl.PostUser(c)
	assertEmptyResponse(http.StatusInternalServerError, t)
}

func TestPostUserBadInput(t *testing.T) {
	setUp()
	stringReader := strings.NewReader("bad-json")
	c.Request = &http.Request{Body: io.NopCloser(stringReader)}
	ctrl.PostUser(c)
	assertEmptyResponse(http.StatusBadRequest, t)
}

func TestGetState(t *testing.T) {
	setUp()
	c.Params = gin.Params{gin.Param{
		Key:   "userid",
		Value: "id",
	}}
	ctrl.GetState(c)
	body := getResponseBody(http.StatusOK, t)
	gameState := presentation.GameState{}
	err := json.Unmarshal(body, &gameState)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if gameState.GamesPlayed != 0 {
		t.Error("failed to assert games played")
	}
	if gameState.Score != 0 {
		t.Error("failed to assert score")
	}
}

func TestGetStateBadInput(t *testing.T) {
	setUp()
	ctrl.GetState(c)
	assertEmptyResponse(http.StatusBadRequest, t)
}

func TestGetStateInternalError(t *testing.T) {
	setUp()
	c.Params = gin.Params{gin.Param{
		Key:   "userid",
		Value: "id",
	}}
	repository.GameStateGetErr = errors.New("error")
	ctrl.GetState(c)
	assertEmptyResponse(http.StatusInternalServerError, t)
}

func TestPutState(t *testing.T) {
	setUp()
	c.Params = gin.Params{gin.Param{
		Key:   "userid",
		Value: "id",
	}}
	stringReader := strings.NewReader("{\"gamesPlayed\": 1, \"score\": 10}")
	c.Request = &http.Request{Body: io.NopCloser(stringReader)}
	ctrl.PutState(c)
	res := getResponseBody(http.StatusOK, t)
	if string(res) != "{}" {
		t.Error("failed to assert empty json response")
	}
}

func TestPutStateBadInput(t *testing.T) {
	setUp()
	ctrl.PutState(c)
	assertEmptyResponse(http.StatusBadRequest, t)
	c.Params = gin.Params{gin.Param{
		Key:   "userid",
		Value: "id",
	}}
	ctrl.PutState(c)
	assertEmptyResponse(http.StatusBadRequest, t)
}

func TestPutStateInternalError(t *testing.T) {
	setUp()
	c.Params = gin.Params{gin.Param{
		Key:   "userid",
		Value: "id",
	}}
	stringReader := strings.NewReader("{\"gamesPlayed\": 1, \"score\": 10}")
	c.Request = &http.Request{Body: io.NopCloser(stringReader)}
	repository.GameStateGetErr = errors.New("error")
	ctrl.PutState(c)
	assertEmptyResponse(http.StatusInternalServerError, t)
}

func TestGetFriends(t *testing.T) {
	setUp()
	c.Params = gin.Params{gin.Param{
		Key:   "userid",
		Value: "id",
	}}
	ctrl.GetFriends(c)
	body := getResponseBody(http.StatusOK, t)
	var friends presentation.Friends
	err := json.Unmarshal(body, &friends)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if len(friends.Friends) != 1 {
		t.Error("failed to assert number of friends")
	} else {
		if friends.Friends[0].Id != "id" {
			t.Error("failed to assert friend ID")
		}
		if friends.Friends[0].Name != "name" {
			t.Error("failed to assert friend name")
		}
		if friends.Friends[0].HighScore != 0 {
			t.Error("failed to assert friend high score")
		}
	}
}

func TestGetFriendsBadInput(t *testing.T) {
	setUp()
	ctrl.GetFriends(c)
	assertEmptyResponse(http.StatusBadRequest, t)
}

func TestGetFriendsInternalError(t *testing.T) {
	setUp()
	c.Params = gin.Params{gin.Param{
		Key:   "userid",
		Value: "id",
	}}
	repository.FriendErr = errors.New("error")
	ctrl.GetFriends(c)
	assertEmptyResponse(http.StatusInternalServerError, t)
}

func TestPutFriends(t *testing.T) {
	setUp()
	c.Params = gin.Params{gin.Param{
		Key:   "userid",
		Value: "id",
	}}
	stringReader := strings.NewReader("{\"friends\":[\"friend-id\"]}")
	c.Request = &http.Request{Body: io.NopCloser(stringReader)}
	ctrl.PutState(c)
	res := getResponseBody(http.StatusOK, t)
	if string(res) != "{}" {
		t.Error("failed to assert empty json response")
	}
}

func TestPutFriendsBadInput(t *testing.T) {
	setUp()
	ctrl.PutFriends(c)
	assertEmptyResponse(http.StatusBadRequest, t)
	c.Params = gin.Params{gin.Param{
		Key:   "userid",
		Value: "id",
	}}
	ctrl.PutFriends(c)
	assertEmptyResponse(http.StatusBadRequest, t)
}

func TestPutFriendsInternalError(t *testing.T) {
	setUp()
	c.Params = gin.Params{gin.Param{
		Key:   "userid",
		Value: "id",
	}}
	stringReader := strings.NewReader("{\"friends\":[\"id\"]}")
	c.Request = &http.Request{Body: io.NopCloser(stringReader)}
	repository.FriendErr = errors.New("error")
	ctrl.PutFriends(c)
	assertEmptyResponse(http.StatusInternalServerError, t)
}

func assertEmptyResponse(code int, t *testing.T) {
	res := getResponseBody(code, t)
	if string(res) != "{}" {
		t.Error("failed to assert that the json response was empty")
	}
}

func getResponseBody(code int, t *testing.T) []byte {
	if code != recorder.Code {
		t.Errorf("failed to assert that code was %d", code)
	}
	body := recorder.Body
	bytes, _ := io.ReadAll(body)
	return bytes
}
