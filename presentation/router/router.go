package router

import (
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
func (s *Server) Routing() {
	s.Router.Use(middleware.Timeout(60 * time.Second))
	s.Router.Use(middleware.Logger)
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
}
