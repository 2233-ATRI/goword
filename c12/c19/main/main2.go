package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.20.253:8888")
	if err != nil {
		fmt.Println("error = ", err)
		return
	}
	fmt.Println("connen", conn)
	rea := bufio.NewReader(os.Stdin) //读取用户的标准输入
	line, err := rea.ReadString('\n')
	if err != nil {
		fmt.Println("reader err is", err)
	}
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("conn write err =", err)
		return
	}
	fmt.Println("发送字节", n)

}
