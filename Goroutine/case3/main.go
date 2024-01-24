package main

import "fmt"

func main() {
	//var wg sync.WaitGroup
	//channel的定义
	ch := make(chan string)

	go func() {
		for {
			msg := <-ch
			fmt.Println(msg)

		}

	}()

	ch <- "1"
	ch <- "2"

}
