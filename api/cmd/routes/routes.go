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
	rtr.POST("/subscriptions", r.h.AddSubscription())
  
	rtr.GET("/friends/list", r.h.FriendsList()) // URL example: 'localhost:3000/friends/list?email=user1@example.com'
	rtr.GET("/friends/common", r.h.CommonFriends()) // URL example: 'localhost:3000/friends/common?email1=user1@example.com&email2=user2@example.com'
}
