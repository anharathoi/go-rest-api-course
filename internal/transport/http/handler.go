package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to comments service
type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Route")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
}
