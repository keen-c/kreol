package server

import (
	"kreol/cmd/web"
	"kreol/internal/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	N := UserHandler{
		user_db: database.UserStore{},
	}
	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/js/*", fileServer)
	r.Get("/",N.GetCreateHandler)
	r.Post("/test", N.Create)
	return r
}

