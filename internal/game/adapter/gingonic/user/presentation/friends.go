package presentation

import "github.com/viktigpetterr/game-rest-api/internal/game/domain"

type JsonFriends struct {
	Friends []friend `json:"friends"`
}

type friend struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	HighScore int    `json:"highscore"`
}

func MakeJsonFriends(friends []domain.Friend) JsonFriends {
	json := JsonFriends{[]friend{}}
	for _, f := range friends {
		json.Friends = append(json.Friends, friend{Id: f.Id, Name: f.Name, HighScore: f.HighScore})
	}

	return json
}
