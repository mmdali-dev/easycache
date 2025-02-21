package async

import (
	"sync"
	"time"
)

type CacheItem struct {
	value      string
	expiration int64
}

type Cache struct {
	items sync.Map
}

func NewAsyncCache(cleanupInterval time.Duration) *Cache {
	cache := &Cache{}

	go func() {
		for {
			time.Sleep(cleanupInterval)
			now := time.Now().UnixNano()
			cache.items.Range(func(key, value interface{}) bool {
				cacheItem := value.(CacheItem)
				if cacheItem.expiration > 0 && now > cacheItem.expiration {
					cache.items.Delete(key)
				}
				return true
			})
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

	c.items.Store(key, CacheItem{
		value:      value,
		expiration: expiration,
	})
}

func (c *Cache) GetWithCheck(key string) (string, bool) {
	item, found := c.items.Load(key)
	if !found {
		return "", false
	}

	cacheItem := item.(CacheItem)
	if cacheItem.expiration > 0 && time.Now().UnixNano() > cacheItem.expiration {
		c.items.Delete(key)
		return "", false
	}

	return cacheItem.value, true
}

func (c *Cache) GetWithoutCheck(key string) (string, bool) {
	item, found := c.items.Load(key)
	if !found {
		return "", false
	}

	cacheItem := item.(CacheItem)
	return cacheItem.value, true
}

func (c *Cache) Delete(key string) {
	c.items.Delete(key)
}

func (c *Cache) Clear() {
	c.items.Range(func(key, _ interface{}) bool {
		c.items.Delete(key)
		return true
	})
}
