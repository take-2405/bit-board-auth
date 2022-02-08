package usecase

import (
	"bit-board-auth/domain/repository"
	"log"
)

type AuthUseCase interface {
	SignUp(userName, email, pass string) (string, error)
	SignIn(email, pass string) (string, error)
}

type authUseCase struct {
	user repository.UserRepository
}

func NewAuthUseCase(user repository.UserRepository) *authUseCase {
	return &authUseCase{user: user}
}

func (uu authUseCase) SignUp(userName, email, pass string) (string, error) {
	uid, err := uu.user.CreateUsersAccount(userName, email, pass)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return uid, nil
}

func (uu authUseCase) SignIn(email, pass string) (string, error) {
	user, err := uu.user.GetUserInfo(email)
	if err != nil {
		log.Println(err)
		return "a", err
	}
	log.Println(user)
	return user.UID, nil
}
