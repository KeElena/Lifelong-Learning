package main

import (
	"fmt"
	"net"
)

func do(conn *net.UDPConn) {
	_, err := conn.Write([]byte("Hello world"))
	if err != nil {
		fmt.Println("Not Internet")
		return
	}
	data := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("404")
		return
	}
	fmt.Println(string(data[:n]))
}

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 7777,
	})
	if err != nil {
		fmt.Println(err)
        return
	}
	defer conn.Close()

	do(conn)
}
