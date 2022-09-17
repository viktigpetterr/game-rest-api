package presentation

import "github.com/viktigpetterr/game-rest-api/internal/game/domain"

type JsonUser struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func MakeJsonUser(user domain.User) JsonUser {
	return JsonUser{
		user.Name,
		user.Id,
	}
}
