package main

import (
	"bit-board-auth/config"
	"bit-board-auth/di"
	"bit-board-auth/infrastructure/disutil"
	router2 "bit-board-auth/presentation/router"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"time"
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

	c := cache.New(5*time.Minute, 10*time.Minute)
	//ルーターを初期化
	router := router2.NewServer()
	//ルーティングとDI
	di.InsertUserDI(router, c)

	//ルーター起動
	port := config.GetServerPort()
	err = http.ListenAndServe(port, router.Router)
	if err != nil {
		log.Fatalf("failed create firebase info file: %v", err)
	}
}
