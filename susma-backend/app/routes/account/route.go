package account

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/josvaal/susma-backend/database"
)

func NewAccountRouter(queries *database.Queries) *chi.Mux {
	accountRouter := chi.NewRouter()

	accountRouter.Post("/account/register", func(w http.ResponseWriter, r *http.Request) {
		registerAccount(w, r, queries)
	})
	accountRouter.Post("/account/login", func(w http.ResponseWriter, r *http.Request) {
		loginAccount(w, r, queries)
	})
	accountRouter.Get("/account/check", func(w http.ResponseWriter, r *http.Request) {
		checkAuth(w, r)
	})
	return accountRouter
}
