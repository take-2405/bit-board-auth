package main

import (
	"bit-board-auth/config"
	router2 "bit-board-auth/presentation/router"
	"net/http"
)

func main() {
	//ルーターを初期化
	router := router2.NewServer()
	//ルーティングとDI
	//di.InsertUserDI(router)
	router.Routing()
	port := config.GetServerPort()
	//ルーター起動
	http.ListenAndServe(port, router.Router)
}
