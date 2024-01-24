package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/quic-go/quic-go"
	"io"
	"net"
)

func main() {
	udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{Port: 1234})
	if err != nil {
		panic(err)
	}
	tr := quic.Transport{
		Conn: udpConn,
	}
	ln, err := tr.Listen(genTlsConfig(), nil)

	go func() {
		for {
			ctx := context.Background()

			conn, err := ln.Accept(ctx)
			if err != nil {
				panic(err)
			}

			go func(ctx context.Context) {
				buf := make([]byte, 1024)
				stream, err := conn.AcceptStream(ctx)
				if err != nil {
					// handle error
				}
				n, err := io.ReadFull(stream, buf)
				if err != nil {
					panic(err)
				}
				// process buffer
				fmt.Printf("Received message: %s\n", string(buf[:n]))
			}(ctx)
		}
	}()
}
func genTlsConfig() *tls.Config {
	cert, err := tls.LoadX509KeyPair("TLS/Cert/cert.pem", "TLS/Cert/key.pem")
	if err != nil {
		panic(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	return config
}
