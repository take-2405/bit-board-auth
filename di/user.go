package di

import (
	"bit-board-auth/infrastructure"
	"bit-board-auth/presentation/controller"
	"bit-board-auth/presentation/router"
	"bit-board-auth/usecase"
	"github.com/patrickmn/go-cache"
	"log"
)

func InsertUserDI(router *router.Server, c *cache.Cache) {
	conn, err := infrastructure.NewFirebase()
	if err != nil {
		log.Fatalf("error %+v\n", err)
	}
	mysqlRepository := infrastructure.NewMysqlRepository()

	userQuery := infrastructure.NewUserPersistence(*mysqlRepository, *conn)
	userUseCase := usecase.NewAuthUseCase(userQuery)
	useHandler := controller.NewUserHandler(userUseCase, c)
	router.Routing(useHandler)
}
