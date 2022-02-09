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
		err := json.NewDecoder(request.Body).Decode(&accountInfo)
		if err != nil {
			log.Println("[ERROR] request bind is err")
			response.RespondError(writer, http.StatusInternalServerError, fmt.Errorf("リクエストの取得に失敗しました"))
		}

		if accountInfo.Email == "" || accountInfo.Pass == "" || accountInfo.UserName == "" {
			log.Println("[ERROR] request bucket is err")
			response.RespondError(writer, http.StatusBadRequest, fmt.Errorf("リクエスト情報が不足しています"))
			return
		}
		if len(accountInfo.Pass) < 6 {
			log.Println("[ERROR] request password is short")
			response.RespondError(writer, http.StatusBadRequest, fmt.Errorf("request password is short"))
			return
		}

		uid, err := uh.userUseCase.SignUp(accountInfo.UserName, accountInfo.Email, accountInfo.Pass)
		if err != nil {
			fmt.Printf("error %+v\n", err)
			response.RespondError(writer, http.StatusInternalServerError, err)
			return
		}

		response.RespondJSON(writer, 200, response.SuccessResponse{Token: uid})
	}
}

func (uh userHandler) SignIn() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var accountInfo request2.CreateUserAccountRequest
		err := json.NewDecoder(request.Body).Decode(&accountInfo)
		if err != nil {
			log.Println("[ERROR] request bind is err")
			response.RespondError(writer, http.StatusInternalServerError, fmt.Errorf("リクエストの取得に失敗しました"))
		}

		if accountInfo.Email == "" || accountInfo.Pass == "" || accountInfo.UserName == "" {
			fmt.Println("[ERROR] request bucket is err")
			response.RespondError(writer, http.StatusBadRequest, fmt.Errorf("リクエスト情報が不足しています"))
			return
		}
		if len(accountInfo.Pass) < 6 {
			fmt.Println("[ERROR] request password is short")
			response.RespondError(writer, http.StatusBadRequest, fmt.Errorf("request password is short"))
			return
		}

		uid, err := uh.userUseCase.SignIn(accountInfo.Email, accountInfo.Pass)
		if err != nil {
			fmt.Printf("error %+v\n", err)
			response.RespondError(writer, http.StatusInternalServerError, err)
			return
		}
		response.RespondJSON(writer, 200, response.SuccessResponse{Token: uid})
	}
}
