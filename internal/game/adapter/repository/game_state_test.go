package repository

import (
	"errors"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql/model"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/repository/testdata"
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
	"testing"
)

func TestNewGameState(t *testing.T) {
	connection := testdata.Connection{}
	repository := GameState{&connection}
	db, _ := connection.Open()

	userModel := model.User{
		Id:   "id",
		Name: "name",
	}
	db.Save(userModel)
	defer db.Delete([]model.User{userModel})

	gameState := domain.GameState{
		UserId:      userModel.Id,
		GamesPlayed: 100, // Should be set to 0
		Score:       50,  // Should be set to 0
	}
	err := repository.New(gameState)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	var gameStateModel model.GameState
	db.Model(model.GameState{UserId: userModel.Id}).Find(&gameStateModel)

	if gameStateModel.GamesPlayed != 0 {
		t.Error("failed to assert games played")
	}
	if gameStateModel.Score != 0 {
		t.Error("failed to assert score")
	}

}

func TestNewGameStateErr(t *testing.T) {
	connection := testdata.Connection{}
	repository := GameState{&connection}
	testdata.ConnectionErr = errors.New("error")
	defer testdata.Reset()
	err := repository.New(domain.GameState{})
	if err != testdata.ConnectionErr {
		t.Error("failed to assert that an error occurred")
	}
}

func TestUpdate(t *testing.T) {
	connection := testdata.Connection{}
	repository := GameState{&connection}
	db, _ := connection.Open()

	userModel := model.User{
		Id:   "id",
		Name: "name",
	}
	db.Save(userModel)
	defer db.Delete([]model.User{userModel})

	gameStateModel := model.GameState{
		UserId:      userModel.Id,
		GamesPlayed: 0,
		Score:       0,
	}
	db.Save(gameStateModel)
	defer db.Delete([]model.GameState{gameStateModel})

	gameState := domain.GameState{
		UserId:      userModel.Id,
		GamesPlayed: 100,
		Score:       50,
	}
	_, err := repository.Update(gameState)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}

	db.Model(model.GameState{UserId: userModel.Id}).Find(&gameStateModel)

	if gameStateModel.GamesPlayed != gameState.GamesPlayed {
		t.Error("failed to assert games played")
	}
	if gameStateModel.Score != gameState.Score {
		t.Error("failed to assert score")
	}
}

func TestUpdateErr(t *testing.T) {
	connection := testdata.Connection{}
	repository := GameState{&connection}
	testdata.ConnectionErr = errors.New("error")
	defer testdata.Reset()
	_, err := repository.Update(domain.GameState{})
	if err != testdata.ConnectionErr {
		t.Error("failed to assert that an error occurred")
	}
}

func TestGetGameStateByUserIdId(t *testing.T) {
	connection := testdata.Connection{}
	repository := GameState{&connection}
	db, _ := connection.Open()

	userModel := model.User{
		Id:   "id",
		Name: "name",
	}
	db.Save(userModel)
	defer db.Delete([]model.User{userModel})

	gameStateModel := model.GameState{
		UserId:      userModel.Id,
		GamesPlayed: 10,
		Score:       5,
	}
	db.Save(gameStateModel)
	defer db.Delete([]model.GameState{gameStateModel})

	gameState, err := repository.GetByUserId(userModel.Id)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if gameStateModel.GamesPlayed != gameState.GamesPlayed {
		t.Error("failed to assert games played")
	}
	if gameStateModel.Score != gameState.Score {
		t.Error("failed to assert score")
	}
}

func TestGetGameStateByUserIdErr(t *testing.T) {
	connection := testdata.Connection{}
	repository := GameState{&connection}
	testdata.ConnectionErr = errors.New("error")
	defer testdata.Reset()
	_, err := repository.GetByUserId("id")
	if err != testdata.ConnectionErr {
		t.Error("failed to assert that an error occurred")
	}
}

func TestGetScoreByUserId(t *testing.T) {
	connection := testdata.Connection{}
	repository := GameState{&connection}
	db, _ := connection.Open()

	userModel := model.User{
		Id:   "id",
		Name: "name",
	}
	db.Save(userModel)
	defer db.Delete([]model.User{userModel})

	gameStateModel := model.GameState{
		UserId:      userModel.Id,
		GamesPlayed: 10,
		Score:       5,
	}
	db.Save(gameStateModel)
	defer db.Delete([]model.GameState{gameStateModel})

	score, err := repository.GetScoreByUserId(userModel.Id)
	if err != nil {
		t.Error("failed to assert that no error occurred")
	}
	if gameStateModel.Score != score {
		t.Error("failed to assert score")
	}
}

func TestGetScoreByUserIdErr(t *testing.T) {
	connection := testdata.Connection{}
	repository := GameState{&connection}
	testdata.ConnectionErr = errors.New("error")
	defer testdata.Reset()
	_, err := repository.GetScoreByUserId("id")
	if err != testdata.ConnectionErr {
		t.Error("failed to assert that an error occurred")
	}
}
