package controller

import (
	middleware2 "bit-board-auth/presentation/middleware"
	request2 "bit-board-auth/presentation/request"
	"bit-board-auth/presentation/response"
	"bit-board-auth/usecase"
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
)

type UserHandler interface {
	SignUp() http.HandlerFunc
	SignIn() http.HandlerFunc
}

type userHandler struct {
	userUseCase usecase.AuthUseCase
	cache       *cache.Cache
}

func NewUserHandler(userUseCase usecase.AuthUseCase, c *cache.Cache) *userHandler {
	return &userHandler{userUseCase: userUseCase, cache: c}
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

		_, ok := uh.cache.Get(accountInfo.Email)
		if ok {
			log.Println("this email already register")
			response.RespondError(writer, http.StatusBadRequest, fmt.Errorf("this email already register"))
			return
		}

		uid, err := uh.userUseCase.SignUp(accountInfo.UserName, accountInfo.Email, accountInfo.Pass)
		if err != nil {
			fmt.Printf("error %+v\n", err)
			response.RespondError(writer, http.StatusInternalServerError, err)
			return
		}

		uh.cache.Set(accountInfo.Email, uid, cache.DefaultExpiration)

		response.RespondJSON(writer, 200, response.SuccessResponse{Token: middleware2.CreateJwt(uid)})
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

		cacheUid, ok := uh.cache.Get(accountInfo.Email)
		if ok {
			log.Println(uh.cache.Get(accountInfo.Email))
			response.RespondJSON(writer, 200, response.SuccessResponse{Token: cacheUid.(string)})
			return
		}

		uid, err := uh.userUseCase.SignIn(accountInfo.Email, accountInfo.Pass)
		if err != nil {
			fmt.Printf("error %+v\n", err)
			response.RespondError(writer, http.StatusInternalServerError, err)
			return
		}

		uh.cache.Set(accountInfo.Email, uid, cache.DefaultExpiration)
		response.RespondJSON(writer, 200, response.SuccessResponse{Token: middleware2.CreateJwt(uid)})
	}
}
