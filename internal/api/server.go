package api

import (
	"net/http"

	"github.com/LYH2263/go-flagship"
)

type Options struct {
	WebDir    string
	AllowCORS bool
}

type Server struct {
	ship *flagship.Ship
	opts Options
	mux  *http.ServeMux
}

func New(ship *flagship.Ship, opts Options) *Server {
	s := &Server{ship: ship, opts: opts, mux: http.NewServeMux()}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.opts.AllowCORS {
		withCORS(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	s.mux.ServeHTTP(w, r)
}
