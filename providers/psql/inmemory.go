package psql

import (
	"context"
	"get-time/model"
	"log"
)

// SaveTimeInMemory implements the providers.TimeStorageInmemory interface.
func (s *Storage) SaveTimeInMemory(_ context.Context, obj model.TimeCanonical) (string, error) {
	s.Lock()
	defer s.Unlock()
	s.inmemory[obj.ID.String()] = obj.Time.String()
	log.Printf("Added time_service inmemory base %v", obj.Time)
	return obj.Time.String(), nil
}
