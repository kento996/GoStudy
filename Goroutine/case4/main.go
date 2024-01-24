package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	go func() {
		fmt.Println("select start...")
		//select具有随机性，当多个case同时满足时，select会随机选择一个执行
		select {
		case msg, OK := <-ch1:
			if !OK {
				fmt.Println("ch1 closed")
				return
			} else {
				fmt.Println("ch1:", msg)
			}
		}
		fmt.Println("select end...")
	}()
	ch1 <- "hello"
	//close(ch1)
	time.Sleep(time.Second * 2)
	fmt.Println("main end...")
}
