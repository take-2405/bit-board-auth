package infrastructure

import (
	"bit-board-auth/domain/repository"
	"firebase.google.com/go/auth"
	"log"
)

type firebasePersistence struct {
	firebase firebaseRepository
}

func NewArticlePersistence(firebase firebaseRepository) repository.UserRepository {
	return &firebasePersistence{firebase: firebase}
}

func (f *firebasePersistence) CreateUsersAccount(userName, email, pass string) error {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(pass).
		DisplayName(userName).
		Disabled(false)
	u, err := f.firebase.Auth.CreateUser(f.firebase.ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %#v\n", u.UserInfo)
	return nil
}

func (f *firebasePersistence) GetUserInfo(id, pass string) error {
	return nil
}
