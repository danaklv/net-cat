package server

import (
	"fmt"
	"time"
)

func (s *Server) broadcastConnectMessage(senderUsername string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	for username, client := range s.clients {
		if username != senderUsername   {
			client.conn.Write([]byte(fmt.Sprintf("\n%s has joined our chat...\n", senderUsername)))
			client.conn.Write([]byte(fmt.Sprintf("[%s][%s]:", currentTime, client.username)))
		}
	}
}

func (s *Server) broadcastDisconnectMessage(senderUsername string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	for username, client := range s.clients {
		if username != senderUsername {
			client.conn.Write([]byte(fmt.Sprintf("\n%s has left our chat...\n", senderUsername)))
			client.conn.Write([]byte(fmt.Sprintf("[%s][%s]:", currentTime, client.username)))
		}
	}
}

func (s *Server) broadcastExceptSender(msg Message, senderUsername string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for username, client := range s.clients {
		if username != senderUsername  {
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			client.conn.Write([]byte(fmt.Sprintf("\n[%s][%s]:%s", currentTime, senderUsername, string(msg.payload))))
			client.conn.Write([]byte(fmt.Sprintf("[%s][%s]:", currentTime, client.username)))
		}
	}
}