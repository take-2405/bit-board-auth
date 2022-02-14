package infrastructure

import "bit-board-auth/domain/repository"

type userRepository struct {
	mysql    mysqlRepository
	firebase firebaseRepository
}

func NewUserPersistence(mysqlConn mysqlRepository, firebaseConn firebaseRepository) repository.UserRepository {
	return &userRepository{mysql: mysqlConn, firebase: firebaseConn}
}
