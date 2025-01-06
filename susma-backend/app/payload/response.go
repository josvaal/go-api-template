package payload

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string          `json:"message"`
	Status  string          `json:"status"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func SendJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error al generar la respuesta JSON", http.StatusInternalServerError)
	}
}

func ChangeResponseData(response *Response, message, status string, data interface{}) {
	if message != "" {
		response.Message = message
	}
	if status != "" {
		response.Status = status
	}
	if data != nil {
		dataBytes, err := json.Marshal(data)
		if err == nil {
			response.Data = json.RawMessage(dataBytes)
		}
	}
}

type ResponseType string

const (
	ErrorServer     ResponseType = "Error del servidor"
	ErrorCommon     ResponseType = "Error"
	ErrorNotFound   ResponseType = "No encontrado"
	ErrorDuplicate  ResponseType = "Duplicado"
	ErrorPermission ResponseType = "Permiso denegado"
	ErrorForbidden  ResponseType = "Prohibido"
	SuccessOk       ResponseType = "Éxito"
	SuccessCreated  ResponseType = "Creado"
)
