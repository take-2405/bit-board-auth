package usecase

import (
	"bit-board-auth/domain/repository"
	"github.com/google/uuid"
	"log"
)

type AuthUseCase interface {
	SignUp(id, pass string) (string, error)
	SignIn(id, pass string) (string, error)
}

type authUseCase struct {
	user repository.UserRepository
}

func NewAuthUseCase(user repository.UserRepository) *authUseCase {
	return &authUseCase{user: user}
}

func (uu authUseCase) SignUp(id, pass string) (string, error) {
	var token string

	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return token, err
	}

	token = uuid.String()
	if err = uu.user.CreateUsersAccount(id, pass, token); err != nil {
		log.Println(err)
		return token, err
	}

	return token, nil
}

func (uu authUseCase) SignIn(id, pass string) (string, error) {
	var token string

	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return token, err
	}

	token = uuid.String()

	if err = uu.user.GetUserInfo(id, pass); err != nil {
		log.Println(err)
		return token, err
	}

	return token, nil
}
