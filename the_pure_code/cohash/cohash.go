package cohash

import (
	"sync"
)

// A Table is a concurrency-safe hash table.
type Table[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

// New returns a new Table.
func New[K comparable, V any]() *Table[K, V] {
	return &Table[K, V]{m: make(map[K]V)}
}

// Set sets the value for key k to v.
func (t *Table[K, V]) Set(k K, v V) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.m[k] = v
}

// Get returns the value for key k, or false if k is not present.
func (t *Table[K, V]) Get(k K) (V, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	v, ok := t.m[k]
	return v, ok
}

// Delete deletes the value for key k.
func (t *Table[K, V]) Delete(k K) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.m, k)
}

// Len returns the number of elements in the table.
func (t *Table[K, V]) Len() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return len(t.m)
}

// Keys returns a slice containing all keys in the table.
func (t *Table[K, V]) Keys() []K {
	t.mu.RLock()
	defer t.mu.RUnlock()
	var keys []K
	for k := range t.m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice containing all values in the table.
func (t *Table[K, V]) Values() []V {
	t.mu.RLock()
	defer t.mu.RUnlock()
	var values []V
	for _, v := range t.m {
		values = append(values, v)
	}
	return values
}

// Clear removes all key-value pairs from the table.
func (t *Table[K, V]) Clear() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.m = make(map[K]V)
}

// Iter calls f on each key-value pair in the table.
func (t *Table[K, V]) Iter(f func(k K, v V)) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for k, v := range t.m {
		f(k, v)
	}
}
