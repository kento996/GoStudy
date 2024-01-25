package main

import (
	"fmt"
	"reflect"
)

// 任意变量可以被赋值为空接口
func main2() {
	a := 1
	i := pr(a)
	fmt.Println(reflect.TypeOf(i))
}
func pr(in interface{}) interface{} {

	return in
}
