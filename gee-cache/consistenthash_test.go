/**
 * @Author QG
 * @Date  2025/1/1 14:48
 * @description
**/

package gee_cache

import (
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	m := New(3, func(data []byte) uint32 {
		key, _ := strconv.Atoi(string(data))
		return uint32(key)
	})

	m.Add("6", "4", "2")

	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testCases {
		if m.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	// Adds 8, 18, 28
	m.Add("8")

	// 27 should now map to 8.
	testCases["27"] = "8"

	for k, v := range testCases {
		if m.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}
}
