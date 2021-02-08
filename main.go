package main

import (
	"fmt"
	"github.com/a2htray/GolangSocketProgramming/client"
	"github.com/a2htray/GolangSocketProgramming/server"
	"log"
	"os"
)

const (
	connHost = "localhost"
	connPort = "4321"
	connType = "tcp"
)

func main() {
	switch os.Args[1] {
	case "server":
		app := server.New(connHost, connPort, connType)
		if err := app.Run(); err != nil {
			log.Fatal(err.Error())
		}
	case "client":
		app := client.New(connHost, connPort, connType)
		if err := app.Run(); err != nil {
			log.Fatal(err.Error())
		}
	default:
		fmt.Println("provide the type [server|client] of application")
	}
}