package infrastructure

import (
	"bit-board-auth/infrastructure/disutil"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

// Driver名
const DriverName = "mysql"

var DBConnectionInfo string

type mysqlRepository struct {
	Client *gorm.DB
}

func NewMysqlRepository() *mysqlRepository {
	var err error
	DBConnectionInfo = disutil.GetMysqlConnectionInfo()
	conn, err := createDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfull DB Connection")
	return &mysqlRepository{Client: conn}
}

func createDBConnection() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(DriverName, DBConnectionInfo)
	if err != nil {
		return nil, err
	}
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
