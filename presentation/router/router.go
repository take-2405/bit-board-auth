package router

import (
	"bit-board-auth/presentation/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	//middleware2 "prac-orm-transaction/presentation/middleware"
	"time"
)

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {
	return &Server{
		Router: chi.NewRouter(),
	}
}

// Router ルーティング設定
func (s *Server) Routing(uh controller.UserHandler) {
	s.Router.Use(middleware.Timeout(60 * time.Second))
	s.Router.Use(middleware.Logger)
	s.Router.Route("/sign", func(api chi.Router) {
		api.Route("/up", func(signup chi.Router) {
			signup.Post("/", uh.SignUp())
		})
		api.Route("/in", func(signup chi.Router) {
			signup.Post("/", uh.SignIn())
		})
	})
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
}
