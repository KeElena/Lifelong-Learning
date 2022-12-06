package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func do(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	info := make([]byte, 128)
	for {

		content, _ := reader.ReadString('\n')
		content = strings.TrimSpace(content)

		if content == "quit" {
			fmt.Println("close conn")
			return
		}

		_, errS := conn.Write([]byte(content))
		if errS != nil {
			fmt.Println("break conn")
			return
		}

		n, errR := conn.Read(info)
		if errR != nil {
			fmt.Println("break conn")
			return
		}
		fmt.Println(string(info[:n]))
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println(err)
		return
	}
	do(conn)
}
