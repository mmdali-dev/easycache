package easycache

import (
	"time"

	"github.com/mmdali-dev/easycache/async"
	"github.com/mmdali-dev/easycache/sync"
)

func NewSyncCache(cleanupInterval time.Duration) *sync.Cache {
	return sync.NewSyncCache(cleanupInterval)
}

func NewAsyncCache(cleanupInterval time.Duration) *async.Cache {
	return async.NewAsyncCache(cleanupInterval)
}
