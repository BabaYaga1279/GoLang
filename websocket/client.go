package websocket

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func clientcreatemsg(addr string, msg string) []byte {
	var arr []byte
	copy(arr, addr+" "+msg)
	return arr
}

func sendmsgtoserver(con net.Conn) {
	localAddr := con.LocalAddr().String()

	for i := 0; i < 4; i++ {
		_, serr := con.Write(clientcreatemsg(localAddr, "Hello."))

		if serr != nil {
			fmt.Printf("Fail sending msg to host: %v\n", serr.Error())
			break
		}

		fmt.Println("msg sent.")
		time.Sleep(time.Millisecond * 500)
	}
}

func recvmsgfromserver(con net.Conn) {
	buff := make([]byte, 1024)

	for i := 0; i < 4; i++ {
		len, rerr := con.Read(buff)

		if rerr != nil {
			fmt.Printf("error reading msg: %v\n", rerr.Error())
			break
		}

		fmt.Printf("msg received: %v\n", string(buff[:len]))
	}
}

func StartClient() {
	fmt.Println("Client running ...")
	ai := NewAddressInfo(1)
	con, err := net.Dial(ai.TYPE, ai.HOST+":"+ai.PORT)

	if err != nil {
		panic(err)
	}

	defer func() {
		con.Close()
		fmt.Println("Connection closed.")
	}()

	fmt.Println("Connected to server.")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		sendmsgtoserver(con)
		wg.Done()
	}()

	go func() {
		recvmsgfromserver(con)
		wg.Done()
	}()

	wg.Wait()
}