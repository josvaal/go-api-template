package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/josvaal/susma-backend/app/models"
	"github.com/josvaal/susma-backend/app/payload"
	"github.com/josvaal/susma-backend/app/utils"
	"github.com/josvaal/susma-backend/database"
)

func registerAccount(w http.ResponseWriter, r *http.Request, queries *database.Queries) {
	ctx := context.Background()
	response := payload.Response{}

	var reqBody models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		payload.ChangeResponseData(&response, "Error al recibir los parámetros", payload.ErrorCommon.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

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

	queries.ResetAutoIncrement(ctx)

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

func loginAccount(w http.ResponseWriter, r *http.Request, queries *database.Queries) {
	ctx := context.Background()
	response := payload.Response{}

	var reqBody models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		payload.ChangeResponseData(&response, "Error al recibir los parámetros", payload.ErrorCommon.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	account, err := queries.GetAccountByEmail(ctx, reqBody.Email)
	if err != nil {
		payload.ChangeResponseData(&response, "Este correo no se encuentra registrado", payload.ErrorNotFound.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusNotFound, response)
		return
	}

	isPassword := utils.CheckPasswordHash(reqBody.Password, account.PasswordHash)

	if !isPassword {
		payload.ChangeResponseData(&response, "Contraseña incorrecta", payload.ErrorCommon.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusUnauthorized, response)
		return
	}

	data := models.Account{
		ID:       account.ID,
		Email:    account.Email,
		Name:     account.FirstName,
		Lastname: account.LastName,
	}

	token, err := utils.GenerateToken(data)
	if err != nil {
		payload.ChangeResponseData(&response, "Error al generar el token", payload.ErrorServer.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	responseData := map[string]interface{}{
		"token": token,
	}

	payload.ChangeResponseData(&response, "Inicio de sesión correcto", nil, responseData)
	payload.SendJSONResponse(w, http.StatusOK, response)
}

func checkAuth(w http.ResponseWriter, r *http.Request) {
	response := payload.Response{}

	token := r.Header.Get("Authorization")
	if token == "" {
		payload.ChangeResponseData(&response, "Error al obtener los permisos", payload.ErrorPermission.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusUnauthorized, response)
		return
	}

	account, err := utils.ValidateToken(token)
	if err != nil {
		payload.ChangeResponseData(&response, "Token inválido o expirado", payload.ErrorPermission.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusUnauthorized, response)
		return
	}

	payload.ChangeResponseData(&response, "Correctamente logueado", nil, account)
	payload.SendJSONResponse(w, http.StatusOK, response)
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	response := payload.Response{}

	token := r.Header.Get("Authorization")
	if token == "" {
		payload.ChangeResponseData(&response, "Error al obtener los permisos", payload.ErrorPermission.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusUnauthorized, response)
		return
	}

	account, err := utils.ValidateToken(token)
	if err != nil {
		payload.ChangeResponseData(&response, "Token inválido o expirado", payload.ErrorPermission.Pointer(), nil)
		payload.SendJSONResponse(w, http.StatusUnauthorized, response)
		return
	}

	payload.ChangeResponseData(&response, "Correctamente logueado", nil, account)
	payload.SendJSONResponse(w, http.StatusOK, response)
}
