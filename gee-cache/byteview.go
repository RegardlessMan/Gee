/**
 * @Author QG
 * @Date  2024/12/29 20:34
 * @description
**/

package gee_cache

//A ByteView holds an immutable view of bytes.

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice.
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String returns the data as a string.
func (v ByteView) String() string {
	return string(v.b)
}

// cloneBytes make a copy of the bytes.
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
