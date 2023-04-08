package providers

import (
	"context"
	"get-time/model"
)

type (
	TimeStoragePsql interface {
		SaveTimePsql(context.Context, model.TimeCanonical) (string, error)
	}
	TimeStorageInmemory interface {
		SaveTimeInMemory(context.Context, model.TimeCanonical) (string, error)
	}
)
