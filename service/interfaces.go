package service

import "context"

type CheckerI interface {
	GetTime(ctx context.Context) (string, error)
}
