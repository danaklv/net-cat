package server

import (
	"fmt"
	"io/ioutil"
)

const maxMessages = 100

func (s *Server) addMessage(msg Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.messages) >= maxMessages {
		s.messages = s.messages[1:]
	}
	s.messages = append(s.messages, msg)
}

func (s *Server) sendChatHistory(client Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, msg := range s.messages {
		client.conn.Write([]byte(fmt.Sprintf("%s%s", msg.from, msg.payload)))
	}
}
func loadWelcomeMessageFromFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
