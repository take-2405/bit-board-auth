package main

import (
	"bit-board-auth/config"
	"bit-board-auth/di"
	"bit-board-auth/infrastructure/disutil"
	router2 "bit-board-auth/presentation/router"
	"log"
	"net/http"
)

func main() {
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatalf("can not read env file: %v", err)
	//}

	//firebaseの接続情報ファイルを作成
	err := disutil.CreateFireBaseConfig()
	if err != nil {
		log.Fatalf("failed create firebase info file: %v", err)
	}

	//ルーターを初期化
	router := router2.NewServer()
	//ルーティングとDI
	di.InsertUserDI(router)

	//ルーター起動
	port := config.GetServerPort()
	err = http.ListenAndServe(port, router.Router)
	if err != nil {
		log.Fatalf("failed create firebase info file: %v", err)
	}
}
