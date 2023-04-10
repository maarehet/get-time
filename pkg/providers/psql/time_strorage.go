package psql

import (
	"context"
	"database/sql"
	"errors"
	"get-time/pkg/model"
	"get-time/pkg/providers"
	"get-time/pkg/providers/psql/schema"
)

var _ providers.TimeStoragePsql = (*Storage)(nil)

func (s *Storage) SaveTimePSQL(ctx context.Context, obj model.TimeCanonical) (string, error) {
	time := schema.NewUserDbModel(obj)
	timeResponse := schema.TimeDb{}
	err := s.db.GetContext(ctx, &timeResponse, "INSERT INTO time_table(id, created_at) VALUES($1,$2)", time.ID, time.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return time.CreatedAt.String(), nil
		}
		s.log.Infof("write error: %v", err)
		return "", err
	}

	return time.CreatedAt.String(), nil
}
