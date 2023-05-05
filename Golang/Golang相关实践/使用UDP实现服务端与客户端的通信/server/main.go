package main

import (
	"fmt"
	"net"
)

func do(conn *net.UDPConn){
	data:=make([]byte,1024)
	n,addr,err:=conn.ReadFromUDP(data)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Print(addr,":")
	fmt.Println(string(data[:n]))
	conn.WriteToUDP([]byte("reply OK"),addr)
}

func main(){
	conn,err:=net.ListenUDP("udp",&net.UDPAddr{
		IP: net.IPv4(127,0,0,1),
		Port: 7777,
	})
	if err !=nil{
		fmt.Println(err)
        return
	}
	defer conn.Close()

	for{
		do(conn)
	}

}
