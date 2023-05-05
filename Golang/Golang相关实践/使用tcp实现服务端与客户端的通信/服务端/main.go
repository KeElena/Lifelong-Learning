package main

import (
	"fmt"
	"net"
)

func do(conn net.Conn) {
	defer conn.Close()
	info := make([]byte, 128)

	for {
		n, err := conn.Read(info)
		if err != nil {
			fmt.Println("break conn")
			return
		}

		_, errS := conn.Write([]byte("state:ok"))

		if errS != nil {
			fmt.Println("break conn")
			return
		}
		fmt.Println(string(info[:n]))
	}
}

func main() {

	Listener, errInit := net.Listen("tcp", "127.0.0.1:7777")
	if errInit != nil {
		fmt.Println(errInit)
		return
	}

	for {
		Conn, errApt := Listener.Accept()

		if errApt != nil {
			fmt.Println(errApt)
			continue
		}
		go do(Conn)
	}
}
