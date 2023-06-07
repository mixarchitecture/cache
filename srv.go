package cache

import (
	"context"
	"time"
)

type Service interface {
	Get(ctx context.Context, k string) (string, error)
	SetEx(ctx context.Context, k string, v interface{}, d time.Duration) error
	Exist(ctx context.Context, k string) (bool, error)
}
