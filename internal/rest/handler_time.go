package rest

import (
	"encoding/json"
	"get-time/internal/rest/model"
	"net/http"
)

func (s *Server) handlerGetTime(w http.ResponseWriter, r *http.Request) {
	timeR := model.TimeResponse{}
	ctx := r.Context()
	time, err := s.service.GetTime(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	timeR.Time = time
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if time != "" {
		if err = json.NewEncoder(w).Encode(timeR); err != nil {
			return
		}
	}

}
