package model

type User struct {
	Id        string    `gorm:"primarykey;uniqueIndex;not null;"`
	Name      string    `gorm:"not null"`
	GameState GameState `gorm:"constraint:OnDelete:CASCADE;"`
	Friends   []*User   `gorm:"many2many:user_friends;constraint:OnDelete:CASCADE;"`
}
