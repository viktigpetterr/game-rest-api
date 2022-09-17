package usecase

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
	"github.com/viktigpetterr/game-rest-api/internal/game/domain/repository"
)

type LoadGameArgs struct {
	UserId string
}

func LoadGameState(gameStateRepository repository.IGameState, args LoadGameArgs) (domain.GameState, error) {
	return gameStateRepository.GetByUserId(args.UserId)
}
