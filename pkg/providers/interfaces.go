package providers

import (
	"context"
	"get-time/pkg/model"
)

type (
	TimeStoragePsql interface {
		SaveTimePSQL(context.Context, model.TimeCanonical) (string, error)
	}
)
