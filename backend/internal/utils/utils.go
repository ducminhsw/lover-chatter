package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string    `json:"message"`
	Code    int       `json:"code"`
	Data    *struct{} `json:"data"`
}

func ParseResponseModel(w http.ResponseWriter, code int, data *struct{}) {
	res := response{
		Message: http.StatusText(code),
		Code:    code,
		Data:    data,
	}
	json.NewEncoder(w).Encode(res)
}
