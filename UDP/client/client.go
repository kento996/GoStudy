package main

import (
	"fmt"
	"net"
)

// UDP 客户端
func main() {
	// 创建UDP socket，net.UDPAddr表示UDP地址
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	//处理错误
	if err != nil {
		fmt.Println("连接UDP服务器失败，err: ", err)
		return
	}
	//最后关闭socket通信
	defer socket.Close()
	// 传输的数据是字节数组
	sendData := []byte("Hello Server")
	// 发送数据
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err: ", err)
		return
	}
	//
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败, err: ", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
