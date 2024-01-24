package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//设置发送端口
	port := flag.String("port", "8360", "port to connect")
	//设置证书文件
	certFile := flag.String("certfile", "TLS/Cert/cert.pem", "trusted CA certificate")
	flag.Parse()
	//读取证书文件
	cert, err := os.ReadFile(*certFile)
	if err != nil {
		log.Fatal(err)
	}
	//将证书文件添加到certPool
	certPool := x509.NewCertPool()
	//判断证书是否添加成功
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		log.Fatalf("unable to parse cert from %s", *certFile)
	}
	//创建tls配置
	config := &tls.Config{RootCAs: certPool}
	//创建tls连接
	conn, err := tls.Dial("tcp", "localhost:"+*port, config)
	if err != nil {
		log.Fatal(err)
	}
	//向conn中发送数据
	_, err = io.WriteString(conn, "Hello simple secure Server\n")
	if err != nil {
		log.Fatal("client write error:", err)
	}
	if err = conn.CloseWrite(); err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 256)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println("client read:", string(buf[:n]))
	conn.Close()
}
