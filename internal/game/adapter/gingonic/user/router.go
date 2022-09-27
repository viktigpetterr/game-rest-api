package user

import (
	"github.com/gin-gonic/gin"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/repository"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/service"
)

func Router(r *gin.Engine) {
	connection := mysql.Connection{}
	ctrl := Controller{
		userRepository:      repository.User{Connection: connection},
		friendRepository:    repository.Friend{Connection: connection},
		gameStateRepository: repository.GameState{Connection: connection},
		uuidService:         service.Uuid{},
	}
	g := r.Group("/user")
	g.GET("", ctrl.GetUsers)
	g.POST("", ctrl.PostUser)
	g.GET("/:userid/state", ctrl.GetState)
	g.PUT("/:userid/state", ctrl.PutState)
	g.GET("/:userid/friends", ctrl.GetFriends)
	g.PUT("/:userid/friends", ctrl.PutFriends)
}
