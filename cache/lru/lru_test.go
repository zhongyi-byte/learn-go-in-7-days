package lru

import (
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

func TestCache_Add(t *testing.T) {
	k1, k2, k3 := "1", "2", "3"
	v1, v2, v3 := "zhongyi", "chuyan", "f"
	cache := New(int64(len(k1+k2+v1+v2)), nil)
	cache.Add(k1, String(v1))
	cache.Add(k2, String(v2))
	if ele, ok := cache.Get(k1); !ok || string(ele.(String)) != "zhongyi" {
		t.Fatalf("cache hit '1' failed")
	}
	cache.Add(k3, String(v3))
	if _, ok := cache.Get(k2); ok || cache.Len() != 2 {
		t.Errorf("%v,%d", ok, cache.Len())
		t.Fatalf("failed to remove oldest")
	}
}
