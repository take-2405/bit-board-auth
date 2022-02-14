package repository

import "firebase.google.com/go/auth"

type UserRepository interface {
	CreateUsersAccountFirebase(userName, email, pass string) (string, error)
	GetUserInfoFirebase(email string) (*auth.UserRecord, error)
	CreateUsersAccountMysql(userId, userName, email, pass string) error
	GetUserInfoMysql(userId string) error
}
