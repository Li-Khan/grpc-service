package http_handler

import (
	"fmt"
	pb "github.com/Li-Khan/grpc-service/api/protobuf/calendar"
	"github.com/Li-Khan/grpc-service/internal/http_handler/middleware"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	client pb.CalendarClient
}

const format string = "2006-01-02"

func NewHandler(client pb.CalendarClient) *http.ServeMux {
	h := &Handler{client: client}

	mux := http.NewServeMux()

	mux.HandleFunc("/add", middleware.PostMiddleware(h.add))
	mux.HandleFunc("/update", middleware.PostMiddleware(h.update))
	mux.HandleFunc("/get", middleware.PostMiddleware(h.get))
	mux.HandleFunc("/list", middleware.PostMiddleware(h.list))
	mux.HandleFunc("/delete", middleware.PostMiddleware(h.delete))

	return mux
}

func (h *Handler) add(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	date := r.URL.Query().Get("date")

	t, err := time.Parse(format, date)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	e, err := h.client.Add(r.Context(), &pb.Event{
		Name: name,
		Date: timestamppb.New(t),
	})
	if err != nil {
		//TODO handle error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, e)
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	paramId := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	date := r.URL.Query().Get("date")

	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t, err := time.Parse(format, date)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e, err := h.client.Update(r.Context(), &pb.Event{
		Id:   id,
		Name: name,
		Date: timestamppb.New(t),
	})
	if err != nil {
		//TODO handle error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, e)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	event, err := h.client.GetByID(r.Context(), &pb.GetEventByIDRequest{Id: id})
	if err != nil {
		//TODO handle error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, event)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	stream, err := h.client.List(r.Context(), &pb.ListEventsRequest{})
	if err != nil {
		//TODO handle error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var events []pb.Event

	for {
		event, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			//TODO handle error
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		events = append(events, *event)
	}

	fmt.Fprint(w, events)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	event, err := h.client.Delete(r.Context(), &pb.DeleteEventRequest{Id: id})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, event)
}
