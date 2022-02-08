package repository

type UserRepository interface {
	CreateUsersAccount(userName, email, pass string) error
	GetUserInfo(id, pass string) error
}
