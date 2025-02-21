package sync

import (
	"time"
)

type CacheItem struct {
	value      string
	expiration int64
}

type Cache struct {
	items map[string]CacheItem
}

func NewSyncCache(cleanupInterval time.Duration) *Cache {
	cache := &Cache{
		items: make(map[string]CacheItem),
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

func (c *Cache) Set(key string, value string, ttl time.Duration) {
	var expiration int64
	if ttl > 0 {
		expiration = time.Now().Add(ttl).UnixNano()
	} else {
		expiration = 0
	}

	c.items[key] = CacheItem{
		value:      value,
		expiration: expiration,
	}
}

func (c *Cache) GetWithCheck(key string) (string, bool) {
	item, found := c.items[key]
	if !found {
		return "", false
	}

	if item.expiration > 0 && time.Now().UnixNano() > item.expiration {
		delete(c.items, key)
		return "", false
	}

	return item.value, true
}

func (c *Cache) GetWithoutCheck(key string) (string, bool) {
	item, found := c.items[key]
	if !found {
		return "", false
	}

	return item.value, true
}

func (c *Cache) Delete(key string) {
	delete(c.items, key)
}

func (c *Cache) Clear() {
	c.items = make(map[string]CacheItem)
}
