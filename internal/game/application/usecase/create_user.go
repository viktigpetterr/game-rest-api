package usecase

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/application/service"
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
	"github.com/viktigpetterr/game-rest-api/internal/game/domain/repository"
)

type CreateUserArgs struct {
	Name string
}

func CreateUser(gameStateRepository repository.IGameState, userRepository repository.IUser, uuidService service.IUuid, args CreateUserArgs) (domain.User, error) {
	user := domain.User{
		Id:   uuidService.New(),
		Name: args.Name,
	}

	err := userRepository.New(user)
	if err != nil {
		return user, err
	}
	user.GameState.UserId = user.Id
	err = gameStateRepository.New(user.GameState)

	return user, err
}
