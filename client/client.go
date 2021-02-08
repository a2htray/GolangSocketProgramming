package client

import (
	"fmt"
	"net"
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
	_, err := net.Dial(c.ConnType, c.ConnHost + ":" + c.ConnPort)
	if err != nil {
		return err
	}
	return nil
}

