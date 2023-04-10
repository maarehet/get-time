package service

import "context"

type Checker interface {
	GetTime(ctx context.Context) (string, error)
}
