package controller

import (
	request2 "bit-board-auth/presentation/request"
	"bit-board-auth/presentation/response"
	"bit-board-auth/usecase"
	"encoding/json"
	"fmt"
	"log"
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

func (uh userHandler) SignUp() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var accountInfo request2.CreateUserAccountRequest
		json.NewDecoder(request.Body).Decode(&accountInfo)

		if accountInfo.Email == "" || accountInfo.Pass == "" || accountInfo.UserName == "" {
			log.Println("[ERROR] request bucket is err")
			response.RespondError(writer, http.StatusBadRequest, fmt.Errorf("リクエスト情報が不足しています"))
			return
		}

		uid, err := uh.userUseCase.SignUp(accountInfo.UserName, accountInfo.Email, accountInfo.Pass)
		if err != nil {
			response.RespondError(writer, http.StatusInternalServerError, err)
			return
		}

		writer.Write([]byte(uid))
	}
}

func (uh userHandler) SignIn() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var accountInfo request2.CreateUserAccountRequest
		json.NewDecoder(request.Body).Decode(&accountInfo)

		if accountInfo.Email == "" || accountInfo.Pass == "" || accountInfo.UserName == "" {
			log.Println("[ERROR] request bucket is err")
			response.RespondError(writer, http.StatusBadRequest, fmt.Errorf("リクエスト情報が不足しています"))
			return
		}
		uid, err := uh.userUseCase.SignIn(accountInfo.Email, accountInfo.Pass)
		if err != nil {
			response.RespondError(writer, http.StatusInternalServerError, err)
			return
		}
		writer.Write([]byte(uid))
	}
}
