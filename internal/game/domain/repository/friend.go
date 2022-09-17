package repository

import "github.com/viktigpetterr/game-rest-api/internal/game/domain"

type IFriend interface {
	PutByUserId(users []domain.User, id string) error
	GetByUserId(id string) ([]domain.Friend, error)
}
