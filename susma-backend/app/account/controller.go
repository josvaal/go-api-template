package account

import (
	"encoding/json"
	"net/http"
)

type AccountResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func registerAccount(w http.ResponseWriter, _ *http.Request) {
	response := AccountResponse{
		Message: "Cuenta registrada con éxito",
		Status:  "success",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error al generar la respuesta JSON", http.StatusInternalServerError)
		return
	}
}

// func loginAccount(w http.ResponseWriter, r *http.Request)
