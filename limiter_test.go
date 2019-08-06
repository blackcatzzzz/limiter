package limiter_test

import (
	"time"

	"github.com/blackcatzzzz/limiter"
	"github.com/blackcatzzzz/limiter/drivers/store/memory"
)

func New(options ...limiter.Option) *limiter.Limiter {
	store := memory.NewStore()
	rate := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  int64(10),
	}
	return limiter.New(store, rate, options...)
}
