package schema

import (
	"time"

	"get-time/pkg/model"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TimeDb struct {
	bun.BaseModel `bun:"time_table,alias:u" db:"omitempty"`
	ID            uuid.UUID `bun:"id" db:"id"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" db:"created_at"`
}

func NewUserDbModel(obj model.TimeCanonical) TimeDb {
	return TimeDb{
		ID:        obj.ID,
		CreatedAt: obj.Time,
	}
}
