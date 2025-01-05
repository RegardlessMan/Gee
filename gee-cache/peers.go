/**
 * @Author QG
 * @Date  2025/1/1 15:02
 * @description
**/

package gee_cache

import pb "Gee/gee-cache/geecachepb"

// PeekPicker PeerPicker is the interface that must be implemented to locate
type PeekPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
