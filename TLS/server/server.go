package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"
)

func main() {
	//指定监听端口
	port := flag.String("port", "8360", "listening port")
	//指定证书文件和私钥文件
	certFile := flag.String("cert", "TLS/Cert/cert.pem", "certificate PEM file")
	keyFile := flag.String("key", "TLS/Cert/key.pem", "key PEM file")
	//解析命令行参数
	flag.Parse()
	//加载证书文件
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}
	//创建TLS配置
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	log.Printf("listening on port %s\n", *port)

	//监听端口,tcp协议+tls协议
	l, err := tls.Listen("tcp", ":"+*port, config)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		//Conn is a generic stream-oriented network connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("accepted connection from %s\n", conn.RemoteAddr())
		//处理连接
		//这里利用goroutine来处理连接，传入的数据是conn，也就是tcp连接
		go func(c net.Conn) {
			//读取数据
			//io.Copy(c, c)
			io.ReadAll(c)
			//关闭连接
			c.Close()
			log.Printf("closing connection from %s\n", conn.RemoteAddr())

		}(conn)
	}
}
