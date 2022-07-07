package http_handler

import (
	"encoding/json"
	"github.com/Li-Khan/grpc-service/internal/helper"
	"log"
	"net/http"

	"google.golang.org/grpc/status"
)

func render(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func errorHandler(w http.ResponseWriter, err error) {
	var code int
	var text string = status.Convert(err).Message()
	log.Println(text)

	switch text {
	case helper.ErrInvalidDate.Error():
		code = http.StatusBadRequest
	case helper.ErrAlreadyExist.Error():
		code = http.StatusConflict
	case helper.ErrEventNotFound.Error():
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}

	w.WriteHeader(code)
}
