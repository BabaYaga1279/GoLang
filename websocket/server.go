package websocket

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func StartServer() {
	fmt.Println("Server running ...")

	ai := NewAddressInfo(0)
	server, err := net.Listen(ai.TYPE, ai.HOST+":"+ai.PORT)

	if err != nil {
		fmt.Printf("error starting server: %v\n", err.Error())
		os.Exit(1)
	}

	defer func() {
		server.Close()
		fmt.Println("Server stopped.")
	}()

	fmt.Println("Listening on " + ai.HOST + ":" + ai.PORT)

	for {
		con, err := server.Accept()
		if err != nil {
			fmt.Printf("error accepting conection: %v\n", err.Error())
			os.Exit(1)
		}

		remoteAddr := con.RemoteAddr().String()
		fmt.Println(remoteAddr + " connected to server.")

		go func() {
			ClientProcessor(con, remoteAddr)
		}()
	}
}

func ClientProcessor(con net.Conn, remoteAddr string) int {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		sendmsgtoclient(con, remoteAddr)
		wg.Done()
	}()

	go func() {
		revcmsgfromclient(con, remoteAddr)
		wg.Done()
	}()

	wg.Wait()

	defer func() {
		fmt.Printf("Finishing seasion with %v and closing connection\n", remoteAddr)
		con.Close()
	}()

	return 0
}

func servercreatemsg(msg string) []byte {
	var arr []byte
	copy(arr, "Server say "+msg)
	return arr
}

func sendmsgtoclient(con net.Conn, remoteAddr string) {
	for {
		_, err := con.Write([]byte{})

		if err != nil {
			fmt.Printf("Error sending msg: %v\n", err.Error())
			return
		}

		fmt.Println("msg sent to " + remoteAddr)
		time.Sleep(time.Microsecond * 1000)
	}
}

func revcmsgfromclient(con net.Conn, remoteAddr string) {
	buff := make([]byte, 1024)

	for {
		len, err := con.Read(buff)

		if err != nil {
			fmt.Printf("Error receiving msg: %v\n", err.Error())
			return
		}

		fmt.Printf("msg received from %v: %v\n", remoteAddr, string(buff[:len]))
	}
}
