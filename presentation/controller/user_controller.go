package controller

import (
	"bit-board-auth/usecase"
	"net/http"
)

type UserHandler interface {
	SignUp() http.HandlerFunc
	SignIn() http.HandlerFunc
}

type userHandler struct {
	userUseCase usecase.AuthUseCase
}

func NewUserHandler(userUseCase usecase.AuthUseCase) *userHandler {
	return &userHandler{userUseCase: userUseCase}
}
