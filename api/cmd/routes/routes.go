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
	rtr.POST("/friends", r.h.AddFriend())
	rtr.GET("/friends/list", r.h.FriendsList())     // URL example: 'localhost:3000/friends/list?email=user1@example.com'
	rtr.POST("/subscriptions/block", r.h.BlockUsers())
}
