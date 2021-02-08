package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Client struct {
	ConnHost string
	ConnPort string
	ConnType string
}

func New(connHost string, connPort string, connType string) *Client {
	return &Client{
		ConnHost: connHost,
		ConnPort: connPort,
		ConnType: connType,
	}
}

func (c *Client)Run() error {
	fmt.Println(fmt.Sprintf("connect to tcp server running on %s:%s", c.ConnHost, c.ConnPort))
	conn, err := net.Dial(c.ConnType, c.ConnHost + ":" + c.ConnPort)
	if err != nil {
		return err
	}

	for {
		fmt.Print("# ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err.Error())
		}

		line = strings.Trim(line, "\n\r")

		if line == "\\q" {
			return nil
		}

		// 向服务器发送数据
		_, err = conn.Write([]byte(line))

		if err != nil {
			fmt.Println(err.Error())
		}

		bytes := make([]byte, 1024)
		length, err := conn.Read(bytes)

		fmt.Println("receive:", string(bytes[0:length]))

	}

}

