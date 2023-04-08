package model

import (
	"time"

	"github.com/google/uuid"
)

type TimeCanonical struct {
	ID   uuid.UUID
	Time time.Time
}

func NewTimeModel() TimeCanonical {
	obj := TimeCanonical{
		ID:   uuid.New(),
		Time: time.Now(),
	}
	return obj
}
