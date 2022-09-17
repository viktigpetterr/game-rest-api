package request

type GameState struct {
	GamesPlayed int `json:"gamesPlayed"`
	Score       int `json:"score"`
}
