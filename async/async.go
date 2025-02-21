package async

import (
	"sync"
	"time"
)

type CacheItem[T any] struct {
	value      T
	expiration int64
}

type Cache[T any] struct {
	items sync.Map
}

func NewAsyncCache[T any](cleanupInterval time.Duration) *Cache[T] {
	cache := &Cache[T]{}

	go func() {
		for {
			time.Sleep(cleanupInterval)
			now := time.Now().UnixNano()
			cache.items.Range(func(key, value interface{}) bool {
				cacheItem, ok := value.(CacheItem[T])
				if !ok {
					return true
				}
				if cacheItem.expiration > 0 && now > cacheItem.expiration {
					cache.items.Delete(key)
				}
				return true
			})
		}
	}()

	return cache
}

func (c *Cache[T]) Set(key string, value T, ttl time.Duration) {
	var expiration int64
	if ttl > 0 {
		expiration = time.Now().Add(ttl).UnixNano()
	} else {
		expiration = 0
	}

	c.items.Store(key, CacheItem[T]{
		value:      value,
		expiration: expiration,
	})
}

func (c *Cache[T]) GetWithCheck(key string) (T, bool) {
	item, found := c.items.Load(key)
	if !found {
		var zero T
		return zero, false
	}

	cacheItem, ok := item.(CacheItem[T])
	if !ok {
		var zero T
		return zero, false
	}

	if cacheItem.expiration > 0 && time.Now().UnixNano() > cacheItem.expiration {
		c.items.Delete(key)
		var zero T
		return zero, false
	}

	return cacheItem.value, true
}

func (c *Cache[T]) GetWithoutCheck(key string) (T, bool) {
	item, found := c.items.Load(key)
	if !found {
		var zero T
		return zero, false
	}

	cacheItem, ok := item.(CacheItem[T])
	if !ok {
		var zero T
		return zero, false
	}

	return cacheItem.value, true
}

func (c *Cache[T]) Delete(key string) {
	c.items.Delete(key)
}

func (c *Cache[T]) Clear() {
	c.items.Range(func(key, _ interface{}) bool {
		c.items.Delete(key)
		return true
	})
}
