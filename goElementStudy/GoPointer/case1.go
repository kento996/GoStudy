package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("a:%d;\na's address:%p\n", a, &a)
	var p *int
	//变量p的类型是一个整型指针，所以只能存放一个整型的地址；
	//*p表示的是指针所指向的整型变量
	//&p表示的是指针的地址
	p = &a
	fmt.Printf("p:%d", *p)
}
