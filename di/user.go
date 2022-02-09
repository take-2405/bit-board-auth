package di

import (
	"bit-board-auth/infrastructure"
	"bit-board-auth/presentation/controller"
	"bit-board-auth/presentation/router"
	"bit-board-auth/usecase"
	"log"
)

func InsertUserDI(router *router.Server) {
	conn, err := infrastructure.NewFirebase()
	if err != nil {
		log.Fatalf("error %+v\n", err)
	}

	userQuery := infrastructure.NewArticlePersistence(*conn)
	userUseCase := usecase.NewAuthUseCase(userQuery)
	useHandler := controller.NewUserHandler(userUseCase)
	router.Routing(useHandler)
}
