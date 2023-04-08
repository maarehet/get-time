package time_service

import (
	"context"

	"get-time/pkg/model"
	"get-time/pkg/providers"
	"get-time/pkg/service"
)

var _ service.Checker = (*TimeService)(nil)

type TimeService struct {
	db       providers.TimeStoragePsql
	inmemory providers.TimeStorageInmemory
}

func (c *TimeService) GetTime(ctx context.Context) (string, error) {
	time := model.NewTimeModel()

	timeNow, err := c.inmemory.SaveTimeInMemory(ctx, time)
	if err != nil {
		return "", err
	}
	timeNow, err = c.inmemory.SaveTimeInMemory(ctx, time)
	if err != nil {
		return "", err // TODO wrap
	}

	return timeNow, nil
}

func NewService(db providers.TimeStoragePsql, inmemory providers.TimeStorageInmemory) *TimeService {
	service := TimeService{
		db:       db,
		inmemory: inmemory,
	}
	return &service
}
