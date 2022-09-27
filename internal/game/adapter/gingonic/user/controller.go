package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/gingonic/user/presentation"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/gingonic/user/request"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/service"
	"github.com/viktigpetterr/game-rest-api/internal/game/application/usecase"
	"github.com/viktigpetterr/game-rest-api/internal/game/domain/repository"
	"gorm.io/gorm"
	"net/http"
)

type Controller struct {
	userRepository      repository.IUser
	friendRepository    repository.IFriend
	gameStateRepository repository.IGameState
	uuidService         service.Uuid
}

func (ctrl Controller) GetUsers(c *gin.Context) {
	users, err := usecase.GetAllUsers(ctrl.userRepository)
	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		return
	}
	c.JSON(http.StatusOK, presentation.MakeUsers(users))
}

func (ctrl Controller) PostUser(c *gin.Context) {
	var req request.User
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	args := usecase.CreateUserArgs{Name: req.Name}
	user, err := usecase.CreateUser(ctrl.gameStateRepository, ctrl.userRepository, ctrl.uuidService, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		return
	}
	c.JSON(http.StatusOK, presentation.MakeUser(user))
}

func (ctrl Controller) GetState(c *gin.Context) {
	userId, isSet := c.Params.Get("userid")
	if !isSet {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}
	args := usecase.LoadGameArgs{UserId: userId}
	gameState, err := usecase.LoadGameState(ctrl.gameStateRepository, args)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, struct{}{})
			return
		}
		c.JSON(http.StatusInternalServerError, struct{}{})
		return
	}

	c.JSON(http.StatusOK, presentation.MakeGameState(gameState))
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
	_, err := usecase.SaveGameState(ctrl.gameStateRepository, args)
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
	friends, err := usecase.GetFriends(ctrl.friendRepository, args)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, struct{}{})
			return
		}
		c.JSON(http.StatusInternalServerError, struct{}{})
		return
	}
	c.JSON(http.StatusOK, presentation.MakeFriends(friends))
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
	err := usecase.UpdateFriends(ctrl.userRepository, ctrl.friendRepository, args)
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
