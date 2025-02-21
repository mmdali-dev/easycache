package easycache

import (
	"time"

	"github.com/mmdali-dev/easycache/async"
	"github.com/mmdali-dev/easycache/sync"
)

func NewSyncCache[T any](cleanupInterval time.Duration) *sync.Cache[T] {
	return sync.NewSyncCache[T](cleanupInterval)
}

func NewAsyncCache[T any](cleanupInterval time.Duration) *async.Cache[T] {
	return async.NewAsyncCache[T](cleanupInterval)
}
