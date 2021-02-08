package server

import (
	"fmt"
	"net"
)

type Server struct {
	ConnHost string
	ConnPort string
	ConnType string
}

func New(connHost string, connPort string, connType string) *Server {
	return &Server{
		ConnHost: connHost,
		ConnPort: connPort,
		ConnType: connType,
	}
}

func (s *Server) Run() error {
	fmt.Println(fmt.Sprintf("%s server running on %s:%s", s.ConnType, s.ConnHost, s.ConnPort))

	listener, err := net.Listen(s.ConnType, s.ConnHost + ":" + s.ConnPort)
	if err != nil {
		return err
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		fmt.Println(fmt.Sprintf("[client %s] connected", conn.RemoteAddr().String()))
	}
}