package repository

import "github.com/viktigpetterr/game-rest-api/internal/game/domain"

type IUser interface {
	New(user domain.User) error
	GetAll() ([]domain.User, error)
	GetByIds(ids []string) ([]domain.User, error)
}
