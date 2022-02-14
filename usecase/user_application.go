package usecase

import (
	"bit-board-auth/domain/repository"
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
	userId, err := uu.user.CreateUsersAccountFirebase(userName, email, pass)
	if err != nil {
		return "", err
	}
	if err = uu.user.CreateUsersAccountMysql(userId, userName, email, pass); err != nil {
		return "", err
	}
	return userId, nil
}

func (uu authUseCase) SignIn(email, pass string) (string, error) {
	user, err := uu.user.GetUserInfoFirebase(email)
	if err != nil {
		return "", err
	}
	return user.UID, nil
}
