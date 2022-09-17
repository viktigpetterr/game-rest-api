package usecase

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
	"github.com/viktigpetterr/game-rest-api/internal/game/domain/repository"
)

type GetFriendsArgs struct {
	UserId string
}

func GetFriends(friendRepository repository.IFriend, args GetFriendsArgs) ([]domain.Friend, error) {
	return friendRepository.GetByUserId(args.UserId)
}
