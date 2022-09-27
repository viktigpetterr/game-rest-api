package presentation

import "github.com/viktigpetterr/game-rest-api/internal/game/domain"

type User struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func MakeUser(user domain.User) User {
	return User{
		user.Name,
		user.Id,
	}
}
