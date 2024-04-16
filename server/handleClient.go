package server

import (
	"fmt"
	"strings"
	"time"
)

func (s *Server) handleClient(client Client) {
	defer func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		s.connections--
		delete(s.clients, client.username)
		client.conn.Close()
		s.broadcastDisconnectMessage(client.username)
	}()

	s.broadcastConnectMessage(client.username)
	s.sendChatHistory(client)

	for {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		client.conn.Write([]byte(fmt.Sprintf("[%s][%s]:", currentTime, client.username)))
		buf := make([]byte, 2048)
		n, err := client.conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.Contains(string(buf[:n]), "[ENTER YOUR NAME]:") {
			continue
		}
		if strings.TrimSpace(string(buf[:n])) == "" {
			continue
		}
		message := Message{
			from:    fmt.Sprintf("[%s][%s]:", currentTime, client.username),
			payload: buf[:n],
		}
		s.addMessage(message)
		s.broadcastExceptSender(message, client.username)
	}
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if s.connections >= 10 {
			fmt.Println("Maximum connections reached")
			conn.Write([]byte("Maximum connections reached. Please try again later.\n"))
			conn.Close()
			continue
		}

		fmt.Println("New connection:", conn.RemoteAddr())
		conn.Write([]byte("Welcome to TCP-Chat!\n"))
		welcomeMsg, err := loadWelcomeMessageFromFile("linux.txt")
		if err != nil {
			fmt.Println("Error loading welcome message:", err)
		} else {
			conn.Write([]byte(welcomeMsg))
		}
		username := ""
		for username == "" {
			conn.Write([]byte("[ENTER YOUR NAME]: "))
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				break
			}
			username = strings.TrimSpace(string(buf[:n]))
			s.mu.Lock()
			_, exists := s.clients[username]
			s.mu.Unlock()
			if exists {
				conn.Write([]byte("Username already taken. Please choose another one.\n"))
				username = ""
			}
		}

		client := Client{conn: conn, username: username}
		s.mu.Lock()
		s.connections++
		s.clients[username] = client
		s.mu.Unlock()
		go s.handleClient(client)
	}
}
