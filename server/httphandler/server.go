package httphandler

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// HTTPHandler is an object representing HTTP's handler(s).
type HTTPHandler struct {
	Mux *mux.Router
}

func InitHTTPHandler() *HTTPHandler {
	return &HTTPHandler{
		Mux: mux.NewRouter().StrictSlash(false),
	}
}

// RegisterRoutes register all routes
func (handlers *HTTPHandler) RegisterRoutes() {
	var routes = handlers.Mux

	routes.Methods("GET").Path("/").HandlerFunc(indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Service is running...")
}
