package api

import (
	"encoding/json"
	"net/http"
)

type ResponseFormat struct {
	StatusCode string      `json:"status_code"`
	Message    string      `json:"status_message"`
	Data       interface{} `json:"data"`
}

type Error struct {
	ErrorMessage string `json:"error_message"`
}

func Response(w http.ResponseWriter, data interface{}, statusCode string, message string, httpCode int) error {

	// Convert the response value to JSON.
	res, err := json.Marshal(ResponseFormat{StatusCode: statusCode, Message: message, Data: data})
	if err != nil {
		return err
	}

	// Respond with the provided JSON.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpCode)
	if _, err := w.Write(res); err != nil {
		return err
	}

	return nil
}

func ResponseOK(w http.ResponseWriter, data interface{}, HTTPStatus int) error {
	return Response(w, data, "200", StatusMessageSuccess, HTTPStatus)
}
