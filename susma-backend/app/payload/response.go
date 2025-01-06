package payload

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string          `json:"message"`
	Error   string          `json:"error,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func SendJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error al generar la respuesta JSON", http.StatusInternalServerError)
	}
}

func ChangeResponseData(response *Response, message string, err *string, data interface{}) {
	if message != "" {
		response.Message = message
	}
	if err != nil {
		response.Error = *err
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
)

func (rt ResponseType) Pointer() *string {
	str := string(rt)
	return &str
}
