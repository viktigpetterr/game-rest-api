package repository

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
)

var UserErr error

type UserMock struct{}

func (_ UserMock) New(_ domain.User) error {
	return UserErr
}

func (_ UserMock) GetAll() ([]domain.User, error) {
	return []domain.User{{
		Id:      "id",
		Name:    "name",
		Friends: nil,
		GameState: domain.GameState{
			UserId:      "id",
			GamesPlayed: 0,
			Score:       0,
		},
	}}, UserErr
}

func (_ UserMock) GetByIds(_ []string) ([]domain.User, error) {
	return nil, UserErr
}
