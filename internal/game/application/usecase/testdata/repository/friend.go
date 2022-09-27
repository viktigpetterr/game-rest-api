package repository

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
)

var FriendErr error

type FriendMock struct{}

func (_ FriendMock) PutByUserId(_ []domain.User, _ string) error {
	return FriendErr
}

func (_ FriendMock) GetByUserId(id string) ([]domain.Friend, error) {
	return []domain.Friend{{
		Id:        id,
		Name:      "name",
		HighScore: 0,
	}}, FriendErr
}
