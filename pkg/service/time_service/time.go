package time_service

import (
	"context"
	"fmt"
	"get-time/pkg/model"
	"get-time/pkg/providers"
	"get-time/pkg/service"
)

var _ service.Checker = (*TimeService)(nil)

type TimeService struct {
	db providers.TimeStoragePsql
}

func (c *TimeService) GetTime(ctx context.Context) (string, error) {
	time := model.NewTimeModel()
	timeNow, err := c.db.SaveTimePSQL(ctx, time)
	if err != nil {
		return "", fmt.Errorf("error saving to the db  %v", err)
	}
	return timeNow, nil
}

func NewService(db providers.TimeStoragePsql) *TimeService {
	s := TimeService{
		db: db,
	}
	return &s
}
