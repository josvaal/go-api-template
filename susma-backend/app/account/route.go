package account

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewAccountRouter(db *sql.DB) *chi.Mux {
	accountRouter := chi.NewRouter()
	accountRouter.Post("/account/register", func(w http.ResponseWriter, r *http.Request) {
		registerAccount(w, r, db)
	})
	return accountRouter
}
