package main

import (
	"log"
	"sync"
)

func demo1() {
	//用于等待协程结束
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println("i:", i)
		}(i)
	}
	wg.Wait()
}
