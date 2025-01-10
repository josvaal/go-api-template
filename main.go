package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/josvaal/susma-backend/app/routes/account"
	"github.com/josvaal/susma-backend/database"
)

func run() (*sql.DB, error) {
	ctx := context.Background()

	db, err := sql.Open("mysql", "root:15022004@/susma?parseTime=true")
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Conexión a la base de datos establecida correctamente")

	queries := database.New(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api", account.NewAccountRouter(queries))

	http.ListenAndServe(":3000", r)
}
