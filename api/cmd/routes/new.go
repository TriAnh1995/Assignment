package routes

import "assignment/internal/handler"

func New(h handler.Handler) Route {
	return Route{h: h}
}
