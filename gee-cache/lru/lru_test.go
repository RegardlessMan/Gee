/**
 * @Author QG
 * @Date  2024/12/28 22:51
 * @description
**/

package lru

import (
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

func TestGet(t *testing.T) {
	cache := New(1024, nil)
	cache.Add("key1", String("1"))
	cache.Add("key2", String("2"))
	cache.Add("key3", String("3"))

	if v, ok := cache.Get("key1"); !ok || string(v.(String)) != "1" {
		t.Fatalf("cache hit key1=1 failed")
	}
}

func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	lru := New(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := []string{}
	onEvicted := func(key string, value Value) {
		keys = append(keys, key)
	}
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	cap := len(k1 + k2 + v1 + v2)

	cache := New(int64(cap), onEvicted)
	cache.Add(k1, String(v1))
	cache.Add(k2, String(v2))
	cache.Add(k3, String(v3))

	if len(keys) != 1 || keys[0] != "key1" {
		t.Fatalf("Call OnEvicted failed")
	}

}
