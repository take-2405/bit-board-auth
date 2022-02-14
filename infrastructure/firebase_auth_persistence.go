package infrastructure

import (
	"firebase.google.com/go/auth"
	"github.com/pkg/errors"
)

//type firebasePersistence struct {
//	firebase firebaseRepository
//}

//func NewArticlePersistence(firebase firebaseRepository) repository.UserRepository {
//	return &firebasePersistence{firebase: firebase}
//}

func (f *userRepository) CreateUsersAccountFirebase(userName, email, pass string) (string, error) {
	_, err := f.firebase.Auth.GetUserByEmail(f.firebase.ctx, email)
	if err == nil {
		return "", errors.New("this email already register")
	}

	params := (&auth.UserToCreate{}).
		Email(email).
		Password(pass).
		DisplayName(userName).
		Disabled(false)

	u, err := f.firebase.Auth.CreateUser(f.firebase.ctx, params)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return u.UID, nil
}

func (f *userRepository) GetUserInfoFirebase(email string) (*auth.UserRecord, error) {
	u, err := f.firebase.Auth.GetUserByEmail(f.firebase.ctx, email)
	if err != nil {
		return nil, errors.New("this email not register")
	}
	return u, nil
}
