package testing

import (
	"github.com/BabaYaga1279/GoLang/websocket"
)

func TestServer() {
	websocket.StartServer()
}

func TestClient() {
	websocket.StartClient()
}
