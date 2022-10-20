package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	DB          struct{}
	Router      *mux.Router
	AllowOrigin string
}

// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
// The handler is typically nil, in which case the DefaultServeMux is used.
// ListenAndServe always returns a non-nil error.
func (s Server) ListenAndServe(port string) error {
	return http.ListenAndServe(":"+port, s.Router)
}

func NewServerSDK() *Server {
	return &Server{}
}
