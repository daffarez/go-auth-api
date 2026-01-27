package handler

import (
	"encoding/json"
	"log"
	"time"
)

type ErrorLog struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Message string `json:"message"`
	Email   string `json:"email,omitempty"`
	Error   string `json:"error"`
}

func LogError(msg string, err error) {
	entry := ErrorLog{
		Time:    time.Now().Format(time.RFC3339),
		Level:   "error",
		Message: msg,
		Error:   err.Error(),
	}
	b, _ := json.Marshal(entry)
	log.Println(string(b))
}
