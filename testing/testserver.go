package testing

import (
	"github.com/BabaYaga1279/GoLang/httpserver"
	"github.com/BabaYaga1279/GoLang/socket"
)

func TestSocketServer() {
	socket.StartServer()
}

func TestSocketClient() {
	socket.StartClient()
}

func TestHTTPServer() {
	httpserver.Test()
}
