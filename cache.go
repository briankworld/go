package main

import (
  "fmt"
  "sync"
  "time"
)

type CacheItem struct {
  Value interface{}
  Expiration int64
}

type Cache struct {
  items map[string]CacheItem
  mu sync.Mutex
}

func NewCache() *Cache {
  return &Cache{
    items: make(map[string]CacheItem)
  }
}

(c *Cache)func set(key string, value interface{}, duration time.Duration) {
  c.mu.Lock()
  defer c.mu.Unlock()

  expiration := time.Now().Add(duration).Unix()
  c.items[key] = CacheItem{
    Value: value,
    Expiration: expiration,
  }
}

(c *Cache)func get(key string) interface{} {
  c.mu.Lock()
  defer c.mu.Unlock()
  item, ok := c.items.get(key)
  if !ok || time.Now().Uinx() > item.Expiration {
    delete(c.items, key)
    return nil
  }

  return item.Value
}

(c *Cache)func delete(key string) {
  c.mu.Lock()
  defer c.mu.Unlock()
  delete(c.items, key)
}
