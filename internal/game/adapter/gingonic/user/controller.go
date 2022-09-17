package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/gingonic/user/presentation"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/gingonic/user/request"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/repository"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/service"
	"github.com/viktigpetterr/game-rest-api/internal/game/application/usecase"
	"gorm.io/gorm"
	"net/http"
)

var (
	userRepository      repository.User
	friendRepository    repository.Friend
	gameStateRepository repository.GameState
	uuidService         service.Uuid
)

type Controller struct{}

func (ctrl Controller) GetUsers(c *gin.Context) {
	users, err := usecase.GetAllUsers(userRepository)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	json := presentation.MakeJsonUsers(users)
	c.JSON(http.StatusOK, json)
}

func (ctrl Controller) PostUser(c *gin.Context) {
	var req request.User
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	args := usecase.CreateUserArgs{Name: req.Name}
	user, err := usecase.CreateUser(gameStateRepository, userRepository, uuidService, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		return
	}

	json := presentation.MakeJsonUser(user)
	c.JSON(http.StatusOK, json)
}

func (ctrl Controller) GetState(c *gin.Context) {
	userId, isSet := c.Params.Get("userid")
	if !isSet {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}
	args := usecase.LoadGameArgs{UserId: userId}
	gameState, err := usecase.LoadGameState(gameStateRepository, args)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, struct{}{})
			return
		}
		c.JSON(http.StatusInternalServerError, struct{}{})
		return
	}

	json := presentation.MakeJsonGameState(gameState)
	c.JSON(http.StatusOK, json)
}

func (ctrl Controller) PutState(c *gin.Context) {
	userId, isSet := c.Params.Get("userid")
	if !isSet {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	var req request.GameState
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	args := usecase.SaveGameArgs{
		UserId:      userId,
		GamesPlayed: req.GamesPlayed,
		Score:       req.Score,
	}
	err := usecase.SaveGameState(gameStateRepository, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		return
	}

	c.JSON(http.StatusOK, struct{}{})
}

func (ctrl Controller) GetFriends(c *gin.Context) {
	userId, isSet := c.Params.Get("userid")
	if !isSet {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	args := usecase.GetFriendsArgs{UserId: userId}
	friends, err := usecase.GetFriends(friendRepository, args)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, struct{}{})
			return
		}
		c.JSON(http.StatusInternalServerError, struct{}{})
		return
	}

	json := presentation.MakeJsonFriends(friends)
	c.JSON(http.StatusOK, json)
}

func (ctrl Controller) PutFriends(c *gin.Context) {
	userId, isSet := c.Params.Get("userid")
	if !isSet {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	var req request.Friends
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	args := usecase.UpdateFriendsArgs{
		UserId:  userId,
		Friends: req.Friends,
	}
	err := usecase.UpdateFriends(userRepository, friendRepository, args)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, struct{}{})
			return
		}
		c.JSON(http.StatusInternalServerError, struct{}{})
		return
	}

	c.JSON(http.StatusOK, struct{}{})
}
