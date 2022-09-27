package presentation

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
	"testing"
)

func TestMakeUsers(t *testing.T) {
	domainFriends := []domain.User{
		{
			Id:   "id-1",
			Name: "name-1",
		},
		{
			Id:   "id-2",
			Name: "name-2",
		},
	}

	users := MakeUsers(domainFriends)
	if len(users.Users) != 2 {
		t.Error("failed to assert number of friends")
	}
}
