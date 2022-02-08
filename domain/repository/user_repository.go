package repository

import "firebase.google.com/go/auth"

type UserRepository interface {
	CreateUsersAccount(userName, email, pass string) (string, error)
	GetUserInfo(email string) (*auth.UserRecord, error)
}
