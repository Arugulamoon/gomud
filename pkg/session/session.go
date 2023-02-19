package session

import (
	"fmt"
	"net"
)

type Session struct {
	Id         string
	Connection net.Conn
}

func (s *Session) SessionId() string {
	return s.Id
}

func (s *Session) WriteLine(str string) error {
	_, err := s.Connection.Write([]byte(str + "\r\n"))
	return err
}

var nextSessionId = 1

func GenerateSessionId() string {
	var sId = nextSessionId
	nextSessionId++
	return fmt.Sprintf("%d", sId)
}
