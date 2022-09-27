package presentation

import "github.com/viktigpetterr/game-rest-api/internal/game/domain"

type Users struct {
	Users []User `json:"users"`
}

func MakeUsers(users []domain.User) Users {
	presentationUsers := []User{}
	for _, user := range users {
		presentationUsers = append(presentationUsers, MakeUser(user))
	}
	return Users{presentationUsers}
}
