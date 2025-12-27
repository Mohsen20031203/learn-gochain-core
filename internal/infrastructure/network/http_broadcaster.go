package network

import (
	"encoding/json"
	"net"
)

type TCPBroadcaster struct {
	peers []string
}

func NewTCPBroadcaster(peers []string) *TCPBroadcaster {
	return &TCPBroadcaster{peers: peers}
}

func (b *TCPBroadcaster) Broadcast(msg Message) {
	for _, peer := range b.peers {
		conn, err := net.Dial("tcp", peer)
		if err != nil {
			continue
		}
		json.NewEncoder(conn).Encode(msg)
		conn.Close()
	}
}
