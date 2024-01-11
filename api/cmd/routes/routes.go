package routes

import (
	"assignment/internal/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	h handler.Handler
}

func (r Route) Routes(rtr *gin.Engine) {
	rtr.POST("/users", r.h.AddUsers())
	rtr.POST("/users/update", r.h.UpdateTopic())

	rtr.POST("/friends", r.h.AddFriend())
	rtr.GET("/friends/list", r.h.FriendsList())
	rtr.GET("/friends/common", r.h.CommonFriends())

	rtr.POST("/subscriptions", r.h.Subscription())
	rtr.POST("/subscriptions/block", r.h.BlockUsers())
}
