package account

import "github.com/go-chi/chi/v5"

func NewAccountRouter() *chi.Mux {
	accountRouter := chi.NewRouter()
	accountRouter.Get("/account/register", registerAccount)
	return accountRouter
}
