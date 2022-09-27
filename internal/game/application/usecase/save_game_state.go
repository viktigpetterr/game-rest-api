package usecase

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
	"github.com/viktigpetterr/game-rest-api/internal/game/domain/repository"
)

type SaveGameArgs struct {
	UserId      string
	GamesPlayed int
	Score       int
}

func SaveGameState(gameStateRepository repository.IGameState, args SaveGameArgs) (domain.GameState, error) {
	oldScore, err := gameStateRepository.GetScoreByUserId(args.UserId)
	if err != nil {
		return domain.GameState{}, err
	}

	if args.Score < oldScore {
		args.Score = oldScore
	}

	gameState := domain.GameState{
		UserId:      args.UserId,
		GamesPlayed: args.GamesPlayed,
		Score:       args.Score,
	}

	return gameStateRepository.Update(gameState)
}
