// broadcaster.go
package network

type Broadcaster interface {
	BroadcastBlock(data []byte) error
	BroadcastTx(data []byte) error
}
