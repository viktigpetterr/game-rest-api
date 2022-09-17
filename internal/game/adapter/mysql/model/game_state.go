package model

type GameState struct {
	UserId      string `gorm:"primarykey"`
	GamesPlayed int
	Score       int
}
