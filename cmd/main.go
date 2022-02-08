package main

import (
	"bit-board-auth/config"
	"bit-board-auth/di"
	router2 "bit-board-auth/presentation/router"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal()
	}
	message := os.Getenv("FS_PRIVATE_KEY")
	log.Println(message)

	//ルーターを初期化
	router := router2.NewServer()
	//ルーティングとDI
	di.InsertUserDI(router)
	port := config.GetServerPort()
	//ルーター起動
	http.ListenAndServe(port, router.Router)
}
