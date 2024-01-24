package main

import (
	"fmt"
	"time"
)

func task() {
	fmt.Println("task start...")
	fmt.Println("task end...")
}
func main() {

	go fmt.Println("aaa")
	//协程之间是并发执行的，所以不保证执行的顺序
	for i := 0; i < 5; i++ {
		go fmt.Println(i)
	}
	//协程调度匿名函数
	go func() {
		fmt.Println("niming")
	}()
	go task()
	time.Sleep(2 * time.Second)

}
