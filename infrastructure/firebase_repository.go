package infrastructure

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

type firebaseRepository struct {
	Auth *auth.Client
	ctx  context.Context
}

func NewFirebase() (*firebaseRepository, error) {
	inst := new(firebaseRepository)
	inst.ctx = context.Background()

	opt := option.WithCredentialsFile("./firebase-auth.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	authInst, err := app.Auth(inst.ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	inst.Auth = authInst
	return inst, nil
}
