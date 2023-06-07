package cache

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
)

type (
	Handler[Entity any] func() (Entity, *i18np.Error)
	Creator[Entity any] func() Entity
)

type Client[Entity any] interface {
	Get(context.Context, string) (Entity, *i18np.Error)
	Handler(Handler[Entity]) Client[Entity]
	Creator(creator Creator[Entity]) Client[Entity]
	Timeout(time.Duration) Client[Entity]
}

func New[Entity any](service Service) Client[Entity] {
	return new[Entity](service)
}
