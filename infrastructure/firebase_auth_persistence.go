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

func (f *firebasePersistence) CreateUsersAccount(userName, email, pass string) (string, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(pass).
		DisplayName(userName).
		Disabled(false)
	u, err := f.firebase.Auth.CreateUser(f.firebase.ctx, params)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Printf("Successfully created user: %#v\n", u.UserInfo)
	return u.UID, nil
}

func (f *firebasePersistence) GetUserInfo(email string) (*auth.UserRecord, error) {
	u, err := f.firebase.Auth.GetUserByEmail(f.firebase.ctx, email)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("Successfully created user: %#v\n", u.UserInfo)
	return u, nil
}
