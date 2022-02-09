package main

import (
	"bit-board-auth/config"
	"bit-board-auth/di"
	router2 "bit-board-auth/presentation/router"
	"net/http"
)

func main() {
	//err := godotenv.Load(".env")
	//if err != nil {
	//	fmt.Printf("読み込み出来ませんでした: %v", err)
	//}

	//ルーターを初期化
	router := router2.NewServer()
	//ルーティングとDI
	di.InsertUserDI(router)
	port := config.GetServerPort()
	//ルーター起動
	http.ListenAndServe(port, router.Router)
}
