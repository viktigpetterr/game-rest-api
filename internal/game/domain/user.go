package domain

type User struct {
	Id        string
	Name      string
	Friends   []User
	GameState GameState
}
