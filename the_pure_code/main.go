package main

import "github.com/appliedgo/ai-times-three/hashtable/cohash"

func main() {
	// Create a new hash table.
	t := cohash.New[string, int]()

	// Set some values.
	t.Set("a", 1)
	t.Set("b", 2)
	t.Set("c", 3)

	// Get a value.
	v, ok := t.Get("b")
	if !ok {
		panic("b not found")
	}
	println(v) // prints "2"

	// Delete a value.
	t.Delete("b")

	// Iter over all key-value pairs.
	t.Iter(func(k string, v int) {
		println(k, v)
	})
}
