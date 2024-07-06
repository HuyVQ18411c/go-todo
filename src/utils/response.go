package utils

import (
	"encoding/json"
	"log"
)

type Response struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetJSONResponse(status int, data any) []byte {
	var stringStatus string
	var message string
	var response []byte

	if status >= 200 && status < 300 {
		stringStatus = "Success"
	} else if status >= 400 && status < 500 {
		stringStatus = "Bad request"
		message = "Invalid request data"
	} else {
		stringStatus = "Failed"
		message = "Internal Server Error. Please try again."
	}

	response, err := json.Marshal(&Response{Status: stringStatus, Data: data, Message: message})
	if err != nil {
		log.Fatal("Failed to parse json data")
		panic(err)
	}

	return response
}
