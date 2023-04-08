package time_service

import (
	"context"
	"get-time/model"
	"get-time/providers"
	"get-time/service"
)

var _ service.CheckerI = (*TimeService)(nil)

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
		return "", err
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
