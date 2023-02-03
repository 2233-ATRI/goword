package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("开始监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("error = ", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端链接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("err = ", err)
		}
		fmt.Println("conn成功", conn)
	}
	fmt.Printf("listen = %v", listen)

}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器的read", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
