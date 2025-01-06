package account

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/josvaal/susma-backend/app/payload"
	"github.com/josvaal/susma-backend/app/utils"
	"github.com/josvaal/susma-backend/database"
)

type RegisterRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	ProfilePicture string `json:"profile_picture"`
}

func registerAccount(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	ctx := context.Background()
	response := payload.Response{}

	var reqBody RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		payload.ChangeResponseData(&response, "Error al recibir los parámetros", payload.ErrorCommon.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	queries := database.New(db)

	_, err := queries.GetAccountByEmail(ctx, reqBody.Email)
	if err == nil {
		payload.ChangeResponseData(&response, "Este correo ya se encuentra registrado", payload.ErrorDuplicate.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusNotFound, response)
		return
	}

	hashPassword, err := utils.HashPassword(reqBody.Password)
	if err != nil {
		payload.ChangeResponseData(&response, "Error al procesar la contraseña", payload.ErrorServer.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	result, err := queries.CreateAccount(ctx, database.CreateAccountParams{
		Email:          reqBody.Email,
		PasswordHash:   hashPassword,
		FirstName:      reqBody.Firstname,
		LastName:       reqBody.Lastname,
		ProfilePicture: reqBody.ProfilePicture,
	})

	if err != nil {
		payload.ChangeResponseData(&response, "Error al crear la cuenta", payload.ErrorServer.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	insertedAuthorID, err := result.LastInsertId()
	if err != nil {
		payload.ChangeResponseData(&response, "Error al obtener la ID de la cuenta creada", payload.ErrorServer.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	data := map[string]interface{}{
		"id": insertedAuthorID,
	}

	payload.ChangeResponseData(&response, "Cuenta creada correctamente", nil, data)
	payload.SendJSONResponse(w, http.StatusCreated, response)
}

// func loginAccount(w http.ResponseWriter, r *http.Request)
