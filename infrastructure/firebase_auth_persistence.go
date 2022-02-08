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

//import (
//"fmt"
//"context"
//
//firebase "firebase.google.com/go"
//"firebase.google.com/go/auth"
//
//"google.golang.org/api/option"
//)
//
//opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
//app, err := firebase.NewApp(context.Background(), nil, opt)
//if err != nil {
//return nil, fmt.Errorf("error initializing app: %v", err)
//}

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
