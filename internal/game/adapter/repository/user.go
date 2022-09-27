package repository

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql/model"
	"github.com/viktigpetterr/game-rest-api/internal/game/domain"
)

type User struct {
	Connection mysql.IConnection
}

func (u User) New(user domain.User) error {
	db, err := u.Connection.Open()
	if err != nil {
		return err
	}
	userModel := model.User{Id: user.Id, Name: user.Name}
	result := db.Create(userModel)
	return result.Error
}

func (u User) GetAll() ([]domain.User, error) {
	db, err := u.Connection.Open()
	if err != nil {
		return nil, err
	}

	var userModels []model.User
	result := db.Find(&userModels)
	err = result.Error
	if err != nil {
		return nil, err
	}

	var users []domain.User
	for _, userModel := range userModels {
		user := domain.User{
			Id:      userModel.Id,
			Name:    userModel.Name,
			Friends: nil,
		}
		users = append(users, user)
	}
	return users, nil
}

func (u User) GetByIds(ids []string) ([]domain.User, error) {
	db, err := u.Connection.Open()
	if err != nil {
		return nil, err
	}
	var userModels []model.User
	err = db.Find(&userModels, ids).Error
	if err != nil {
		return nil, err
	}

	var users []domain.User
	for _, userModel := range userModels {
		user := domain.User{
			Id:      userModel.Id,
			Name:    userModel.Name,
			Friends: nil,
		}
		users = append(users, user)
	}
	return users, nil
}
