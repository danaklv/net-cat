package server

import "net"

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAdd: listenAddr,
		quitch:    make(chan struct{}),
		clients:   make(map[string]Client),
		messages:  make([]Message, 0, maxMessages),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAdd)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch

	return nil
}
