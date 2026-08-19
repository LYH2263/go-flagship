package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/LYH2263/go-flagship"
)

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, err error) {
	code := http.StatusBadRequest
	switch {
	case errors.Is(err, flagship.ErrNotFound):
		code = http.StatusNotFound
	case errors.Is(err, flagship.ErrClosed), errors.Is(err, flagship.ErrNilStore):
		code = http.StatusServiceUnavailable
	}
	writeJSON(w, code, map[string]string{"error": err.Error()})
}
