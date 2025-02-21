package sync

import (
	"time"
)

type CacheItem[T any] struct {
	value      T
	expiration int64
}

type Cache[T any] struct {
	items map[string]CacheItem[T]
}

func NewSyncCache[T any](cleanupInterval time.Duration) *Cache[T] {
	cache := &Cache[T]{
		items: make(map[string]CacheItem[T]),
	}

	go func() {
		for {
			time.Sleep(cleanupInterval)
			now := time.Now().UnixNano()
			for key, item := range cache.items {
				if item.expiration > 0 && now > item.expiration {
					delete(cache.items, key)
				}
			}
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

	c.items[key] = CacheItem[T]{
		value:      value,
		expiration: expiration,
	}
}

func (c *Cache[T]) GetWithCheck(key string) (T, bool) {
	var zero T
	item, found := c.items[key]
	if !found {
		return zero, false
	}

	if item.expiration > 0 && time.Now().UnixNano() > item.expiration {
		delete(c.items, key)
		return zero, false
	}

	return item.value, true
}

func (c *Cache[T]) GetWithoutCheck(key string) (T, bool) {
	var zero T
	item, found := c.items[key]
	if !found {
		return zero, false
	}

	return item.value, true
}

func (c *Cache[T]) Delete(key string) {
	delete(c.items, key)
}

func (c *Cache[T]) Clear() {
	c.items = make(map[string]CacheItem[T])
}
