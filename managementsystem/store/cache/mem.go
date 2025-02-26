package cache

import (
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/phannita016/management/store"
)

type (
	// MemOption memory cache optional type.
	MemOption func(*MemCache)

	// MemCache memory cache struct.
	MemCache struct {
		TTL  time.Duration
		Size int
	}
)

// MemTTL set time to live or expire time in cache.
func MemTTL(ttl time.Duration) MemOption {
	return func(m *MemCache) {
		m.TTL = ttl
	}
}

// MemSize set size in memory cache.
func MemSize(size int) MemOption {
	return func(m *MemCache) {
		m.Size = size
	}
}

type mem[T any] struct {
	lru *expirable.LRU[string, T]
}

// NewMemCache function input optional memory cache.
func NewMemCache[T any](opts ...MemOption) store.Cache[T] {
	// default memory cache.
	mc := MemCache{TTL: time.Minute, Size: 100}

	for _, opt := range opts {
		opt(&mc)
	}

	// return memory cache set size and ttl.
	return &mem[T]{
		lru: expirable.NewLRU[string, T](mc.Size, nil, mc.TTL),
	}
}

// NewMem memory cache implement lru return interface.
func NewMem[T any](lru *expirable.LRU[string, T]) store.Cache[T] {
	return &mem[T]{lru: lru}
}

// Get function input argument key string return generic type, bool
func (m *mem[T]) Get(key string) (T, bool) {
	return m.lru.Get(key)
}

// Set function set key, t generic type.
func (m *mem[T]) Set(key string, t T) bool {
	return m.lru.Add(key, t)
}

// Delete function key from the cache
func (m *mem[T]) Delete(key string) bool {
	return m.lru.Remove(key)
}

// Keys function return keys slices string.
func (m *mem[T]) Keys() []string {
	return m.lru.Keys()
}

// Values return slices generic type values.
func (m *mem[T]) Values() []T {
	return m.lru.Values()
}
