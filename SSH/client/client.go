package main

import (
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {

	// 设置 SSH 客户端配置
	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 注意：在生产环境中应使用更安全的方式
	}

	// 连接到 SSH 服务器
	_, err := ssh.Dial("tcp", "remotehost:22", config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}

}
