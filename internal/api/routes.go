package api

import "net/http"

func (s *Server) routes() {
	s.mux.HandleFunc("/api/health", s.handleHealth)
	s.mux.HandleFunc("/api/flags", s.handleFlags)
	s.mux.HandleFunc("/api/flags/", s.handleFlagOne)
	s.mux.HandleFunc("/api/eval", s.handleEval)
	s.mux.HandleFunc("/api/stats", s.handleStats)
	if s.opts.WebDir != "" {
		s.mux.Handle("/", staticHandler(s.opts.WebDir))
	} else {
		s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
		})
	}
}
