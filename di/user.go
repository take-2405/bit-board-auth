package di

import (
	"bit-board-auth/infrastructure"
	"bit-board-auth/presentation/controller"
	"bit-board-auth/presentation/router"
	"bit-board-auth/usecase"
)

func InsertUserDI(router *router.Server) {
	conn := infrastructure.NewFirebase()

	userQuery := infrastructure.NewArticlePersistence(*conn)

	userUseCase := usecase.NewAuthUseCase(userQuery)

	useHandler := controller.NewUserHandler(userUseCase)
	router.Routing(useHandler)
}
