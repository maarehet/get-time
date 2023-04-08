package psql

import (
	"context"
	"fmt"
	"get-time/model"
	"get-time/providers"
	"get-time/providers/psql/schema"
)

var _ providers.TimeStoragePsql = (*Storage)(nil)

func (s *Storage) SaveTimePsql(ctx context.Context, obj model.TimeCanonical) (string, error) {
	time := schema.NewUserDbModel(obj)

	timeResponse := schema.TimeDb{}
	err := s.db.GetContext(ctx, &timeResponse, "INSERT INTO time_table(id, created_at) VALUES($1,$2)", time.ID, time.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return "", err

	}

	return timeResponse.CreatedAt.String(), nil
}
