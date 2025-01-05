package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"

	"github.com/josvaal/susma-backend/database"
)

func run() error {
	ctx := context.Background()

	db, err := sql.Open("mysql", "root:15022004@/susma?parseTime=true")
	if err != nil {
		return err
	}

	queries := database.New(db)

	// list all authors
	authors, err := queries.ListAccounts(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	result, err := queries.CreateAccount(ctx, database.CreateAccountParams{
		Email:          "josval.personal@gmail.com",
		PasswordHash:   "15022004",
		FirstName:      "José",
		LastName:       "Valentino",
		ProfilePicture: "",
	})
	if err != nil {
		return err
	}

	insertedAccountID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(insertedAccountID)

	// get the author we just inserted
	fetchedAccount, err := queries.GetAccount(ctx, insertedAccountID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAccountID, fetchedAccount.ID))

	err_q := queries.CreateSubscription(ctx, database.CreateSubscriptionParams{
		AccountID:        int32(fetchedAccount.ID),
		ServiceName:      "nesflis",
		PlanName:         "xd",
		BillingFrequency: "monthly",
		Cost:             10.50,
		Currency:         "S/. ",
		Icon:             "xd",
	})
	if err_q != nil {
		return err_q
	}

	suscription, err := queries.GetSubscription(ctx, 1)
	if err != nil {
		return err
	}
	log.Println(suscription)

	return nil

}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
