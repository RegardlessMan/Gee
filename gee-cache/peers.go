/**
 * @Author QG
 * @Date  2025/1/1 15:02
 * @description
**/

package gee_cache

// PeekPicker PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key.
type PeekPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
