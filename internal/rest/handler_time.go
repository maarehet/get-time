package rest

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handlerGetTime(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	time, err := s.service.GetTime(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if time != "" {
		if err = json.NewEncoder(w).Encode(time); err != nil {
			return
		}
	}
}
