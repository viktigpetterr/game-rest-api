package usecase

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/domain/repository"
)

type UpdateFriendsArgs struct {
	UserId  string
	Friends []string
}

func UpdateFriends(userRepository repository.IUser, friendRepository repository.IFriend, args UpdateFriendsArgs) error {
	users, err := userRepository.GetByIds(args.Friends)
	if err != nil {
		return err
	}
	return friendRepository.PutByUserId(users, args.UserId)
}
