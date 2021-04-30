package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen error", err.Error())
		return
	}
	for  {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("acc err", err.Error())
			continue
		}
		go Process(conn)
	}
}

func Process(conn net.Conn) {
	defer conn.Close()
	//for {
	//	reader := bufio.NewReader(conn)
	//	var buf [128]byte
	//	//n, err := reader.Read(buf[:])
	//	if err != nil {
	//		break
	//	}
	//	//rec := string[bu]
	//}
}

