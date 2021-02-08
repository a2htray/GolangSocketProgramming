package server

import (
	"fmt"
	"github.com/a2htray/GolangSocketProgramming/operator"
	"net"
	"strconv"
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

		go receiveMessage(conn)
	}
}

func receiveMessage(conn net.Conn) {
	for {
		bytes := make([]byte, 1024)
		length, err := conn.Read(bytes)

		if err != nil {
			conn.Write([]byte(err.Error()))
			continue
		}

		expression := string(bytes[0:length])
		fmt.Println(fmt.Sprintf(
			"receive instruction from [client %s]: %s",
			conn.RemoteAddr().String(), expression))

		operatorFunc, a, b, err := operator.ParseExpression(expression)

		if err != nil {
			conn.Write([]byte(err.Error()))
			continue
		}

		conn.Write([]byte(strconv.FormatFloat(operatorFunc(a, b), 'f', 10, 64)))
	}
}
