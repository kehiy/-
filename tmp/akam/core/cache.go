package core

import (
	"sync"
	"time"
)

type Cache struct {
	lock sync.RWMutex
	Data map[string][]byte
}

func New() *Cache {
	return &Cache{
		Data: make(map[string][]byte),
	}
}

func (c *Cache) Delete(key []byte) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.Data, string(key))
	return true
}

func (c *Cache) Have(key []byte) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	_, ok := c.Data[string(key)]
	return ok
}

func (c *Cache) Get(key []byte) ([]byte, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	val, ok := c.Data[string(key)]
	if !ok {
		return nil, ErrorGet
	}
	return val, nil
}

func (c *Cache) Set(key, value []byte, ttl time.Duration) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.Data[string(key)] = value

	if ttl > 0 {
		go func() {
			<-time.After(ttl)
			delete(c.Data, string(key))
		}()
	}

	return nil
}
