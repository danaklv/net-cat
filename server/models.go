package server

import (
	"net"
	"sync"
)

type Client struct {
	conn     net.Conn
	username string
}

type Message struct {
	from    string
	payload []byte
}
type ClientExit struct {
	Username string
}
type Server struct {
	listenAdd   string
	ln          net.Listener
	quitch      chan struct{}
	clients     map[string]Client
	messages    []Message
	mu          sync.Mutex
	connections int
}
