package infrastructure

import (
	"bit-board-auth/infrastructure/table"
	"github.com/pkg/errors"
)

func (u *userRepository) CreateUsersAccountMysql(userId, userName, email, pass string) error {
	usersInfo := table.Users{ID: userId, Password: pass, Email: email, Name: userName}
	var dataExistsCheck table.Users

	u.mysql.Client.First(&dataExistsCheck, "id=?", usersInfo.ID)
	if dataExistsCheck.ID != "" {
		return errors.New("this userID is already registered")
	}

	if err := u.mysql.Client.Create(&usersInfo).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (u *userRepository) GetUserInfoMysql(userId string) error {
	return nil
}
