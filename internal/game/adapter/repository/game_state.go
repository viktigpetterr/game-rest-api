package repository

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql/model"
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
)

type GameState struct{}

func (_ GameState) New(gameState domain.GameState) error {
	db, err := mysql.Connection()
	if err != nil {
		return err
	}

	err = db.
		Create(
			model.GameState{
				UserId:      gameState.UserId,
				GamesPlayed: gameState.GamesPlayed,
				Score:       gameState.Score,
			}).
		Error

	return err
}

func (_ GameState) Update(gameState domain.GameState) error {
	db, err := mysql.Connection()
	if err != nil {
		return err
	}

	gameStateModel := model.GameState{
		GamesPlayed: gameState.GamesPlayed,
		Score:       gameState.Score,
	}
	result := db.Model(&model.GameState{UserId: gameState.UserId}).
		Updates(gameStateModel)

	return result.Error
}

func (_ GameState) GetByUserId(id string) (domain.GameState, error) {
	db, err := mysql.Connection()
	if err != nil {
		return domain.GameState{}, err
	}

	var gameStateModel model.GameState
	err = db.First(&gameStateModel, "user_id = ?", id).Error
	if err != nil {
		return domain.GameState{}, err
	}

	gameState := domain.GameState{
		UserId:      gameStateModel.UserId,
		GamesPlayed: gameStateModel.GamesPlayed,
		Score:       gameStateModel.Score,
	}
	return gameState, nil
}

func (_ GameState) GetScoreByUserId(id string) (int, error) {
	db, err := mysql.Connection()
	if err != nil {
		return 0, err
	}

	gameStateModel := model.GameState{UserId: id}
	err = db.Select("score").Find(&gameStateModel).Error
	return gameStateModel.Score, err
}
