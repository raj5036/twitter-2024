package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ResponseFormat struct {
	StatusCode string      `json:"status_code"`
	Message    string      `json:"status_message"`
	Data       interface{} `json:"data"`
}

type ErrorFormat struct {
	StatusCode   string      `json:"status_code"`
	ErrorMessage string      `json:"error_message"`
	Error        interface{} `json:"error"`
}

func Response(w http.ResponseWriter, res []byte, err error, httpCode int) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpCode)
	if _, err := w.Write(res); err != nil {
		return err
	}

	return nil
}

func ResponseOK(w http.ResponseWriter, data interface{}, httpCode int) error {
	response := ResponseFormat{
		StatusCode: strconv.Itoa(httpCode),
		Message:    StatusMessageSuccess,
		Data:       data,
	}
	res, err := json.Marshal(response)
	return Response(w, res, err, httpCode)
}

func ResponseError(w http.ResponseWriter, message string, httpCode int) error {
	error := ErrorFormat{
		Error:        nil,
		StatusCode:   strconv.Itoa(httpCode),
		ErrorMessage: message,
	}

	res, err := json.Marshal(error)
	return Response(w, res, err, httpCode)
}
