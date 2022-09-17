package presentation

import "github.com/viktigpetterr/game-rest-api/internal/game/domain"

type JsonUsers struct {
	Users []JsonUser `json:"users"`
}

func MakeJsonUsers(users []domain.User) JsonUsers {
	jsonUsers := []JsonUser{}
	for _, user := range users {
		jsonUsers = append(jsonUsers, MakeJsonUser(user))
	}
	return JsonUsers{jsonUsers}
}
