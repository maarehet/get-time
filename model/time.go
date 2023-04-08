package model

import (
	"github.com/google/uuid"
	"time"
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
