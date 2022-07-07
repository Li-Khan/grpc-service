package http_handler

import (
	pb "github.com/Li-Khan/grpc-service/api/protobuf/calendar"
	"net/http"
)

type Handler struct {
	client pb.CalendarClient
}

func NewHandler(client pb.CalendarClient) *http.ServeMux {
	h := &Handler{client: client}

	mux := http.NewServeMux()
	mux.HandleFunc("/test", h.test)

	return mux
}

func (h *Handler) test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}
