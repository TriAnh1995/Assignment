package handler

import (
	"assignment/internal/controller"
)

type Handler struct {
	ctrl controller.Controller
}

func New(ctrl controller.Controller) Handler {
	return Handler{ctrl}
}
