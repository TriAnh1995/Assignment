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
}
