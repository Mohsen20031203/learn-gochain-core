package network

import (
	"encoding/json"
	"net"
)

type Message struct {
	Type string `json:"type"` // "block" | "tx"
	Data []byte `json:"data"`
}

type TCPServer struct {
	address string
	handler func(Message)
}

func NewTCPServer(addr string, handler func(Message)) *TCPServer {
	return &TCPServer{
		address: addr,
		handler: handler,
	}
}

func (s *TCPServer) Start() error {
	ln, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				continue
			}
			go s.handleConn(conn)
		}
	}()
	return nil
}

func (s *TCPServer) handleConn(conn net.Conn) {
	defer conn.Close()

	var msg Message
	if err := json.NewDecoder(conn).Decode(&msg); err != nil {
		return
	}

	s.handler(msg)
}
