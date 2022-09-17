package usecase

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
	"github.com/viktigpetterr/game-rest-api/internal/game/domain/repository"
)

func GetAllUsers(userRepository repository.IUser) ([]domain.User, error) {
	return userRepository.GetAll()
}
