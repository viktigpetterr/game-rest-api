package presentation

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
	"testing"
)

func TestMakeUser(t *testing.T) {
	domainUser := domain.User{
		Name: "name",
		Id:   "id",
	}

	user := MakeUser(domainUser)
	if user.Name != domainUser.Name {
		t.Error("failed to assert games played")
	}
	if user.Id != domainUser.Id {
		t.Error("failed to assert score")
	}
}
