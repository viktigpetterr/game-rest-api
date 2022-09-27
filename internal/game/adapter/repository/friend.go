package repository

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql/model"
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
)

type Friend struct {
	Connection mysql.IConnection
}

func (f Friend) PutByUserId(users []domain.User, id string) error {
	db, err := f.Connection.Open()
	if err != nil {
		return err
	}
	var friends []model.User
	for _, user := range users {
		friends = append(friends, model.User{Id: user.Id})
	}

	err = db.
		Model(&model.User{Id: id}).
		Omit("User.id").
		Association("Friends").
		Replace(friends)
	return err
}

func (f Friend) GetByUserId(id string) ([]domain.Friend, error) {
	db, err := f.Connection.Open()
	if err != nil {
		return nil, err
	}

	var friends []domain.Friend
	err = db.
		Table("user_friends").
		Select("user_friends.friend_id as id, users.name as name, game_states.score as highscore").
		Joins("inner join users on users.id = user_friends.friend_id").
		Joins("inner join game_states on game_states.user_id = user_friends.friend_id").
		Where("user_friends.user_id = ?", id).
		Scan(&friends).Error
	return friends, err
}
