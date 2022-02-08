package infrastructure

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type firebaseRepository struct {
	Auth *auth.Client
	ctx  context.Context
}

func NewFirebase() *firebaseRepository {
	inst := new(firebaseRepository)
	inst.ctx = context.Background()

	opt := option.WithCredentialsFile("./.secret/bit-board-dcf2c-firebase-adminsdk-rudoq-a5211118ec.json")
	// GOOGLE_APPLICATION_CREDENTIALSで指定した認証情報ファイルを暗黙的に読み込む
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return inst
	}
	authInst, err := app.Auth(inst.ctx)
	if err != nil {
		return inst
	}

	inst.Auth = authInst
	return inst
}
