package presentation

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
	"testing"
)

func TestMakeFriends(t *testing.T) {
	domainFriends := []domain.Friend{
		{
			Id:        "id-1",
			Name:      "name-1",
			HighScore: 0,
		},
		{
			Id:        "id-2",
			Name:      "name-2",
			HighScore: 0,
		},
	}

	friends := MakeFriends(domainFriends)
	if len(friends.Friends) != 2 {
		t.Error("failed to assert number of friends")
	}
}
