package cohash

import (
	"testing"
)

// TestTable tests the Table type.
func TestTable(t *testing.T) {
	t.Run("Set", func(t *testing.T) {
		t.Parallel()
		tbl := New[string, int]()
		tbl.Set("a", 1)
		if v, ok := tbl.Get("a"); !ok || v != 1 {
			t.Fatal("expected a=1")
		}
	})

	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		tbl := New[string, int]()
		if _, ok := tbl.Get("a"); ok {
			t.Fatal("expected a not found")
		}
		tbl.Set("a", 1)
		if v, ok := tbl.Get("a"); !ok || v != 1 {
			t.Fatal("expected a=1")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		t.Parallel()
		tbl := New[string, int]()
		tbl.Set("a", 1)
		tbl.Delete("a")
		if _, ok := tbl.Get("a"); ok {
			t.Fatal("expected a not found")
		}
	})

	t.Run("Len", func(t *testing.T) {
		t.Parallel()
		tbl := New[string, int]()
		if tbl.Len() != 0 {
			t.Fatal("expected empty table")
		}
		tbl.Set("a", 1)
		if tbl.Len() != 1 {
			t.Fatal("expected len=1")
		}
	})

	t.Run("Keys", func(t *testing.T) {
		t.Parallel()
		tbl := New[string, int]()
		tbl.Set("a", 1)
		tbl.Set("b", 2)
		tbl.Set("c", 3)
		keys := tbl.Keys()
		if len(keys) != 3 {
			t.Fatal("expected 3 keys")
		}
		for _, k := range keys {
			if _, ok := tbl.Get(k); !ok {
				t.Fatal("expected key to be present")
			}
		}
	})

	t.Run("Values", func(t *testing.T) {
		t.Parallel()
		tbl := New[string, int]()
		tbl.Set("a", 1)
		tbl.Set("b", 2)
		tbl.Set("c", 3)
		values := tbl.Values()
		if len(values) != 3 {
			t.Fatal("expected 3 values")
		}
		keys := tbl.Keys()
		for _, v := range values {
			found := false
			for _, k := range keys {
				if val, ok := tbl.Get(k); ok && val == v {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("expected value %d to be present", v)
			}
		}
	})

	t.Run("Clear", func(t *testing.T) {
		t.Parallel()
		tbl := New[string, int]()
		tbl.Set("a", 1)
		tbl.Set("b", 2)
		tbl.Set("c", 3)
		tbl.Clear()
		if tbl.Len() != 0 {
			t.Fatal("expected empty table")
		}
	})

	t.Run("Iter", func(t *testing.T) {
		t.Parallel()
		tbl := New[string, int]()
		tbl.Set("a", 1)
		tbl.Set("b", 2)
		tbl.Set("c", 3)
		var keys []string
		tbl.Iter(func(k string, v int) {
			keys = append(keys, k)
		})
		if len(keys) != 3 {
			t.Fatal("expected 3 keys")
		}
		for _, k := range keys {
			if _, ok := tbl.Get(k); !ok {
				t.Fatal("expected key to be present")
			}
		}
	})
}
