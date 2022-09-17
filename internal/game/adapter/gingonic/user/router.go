package user

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	ctrl := Controller{}
	g := r.Group("/user")
	g.GET("", ctrl.GetUsers)
	g.POST("", ctrl.PostUser)
	g.GET("/:userid/state", ctrl.GetState)
	g.PUT("/:userid/state", ctrl.PutState)
	g.GET("/:userid/friends", ctrl.GetFriends)
	g.PUT("/:userid/friends", ctrl.PutFriends)
}
