package main

import (
	"fmt"
	"strings"
)

func main() {
	//判断string的开头和结尾
	str := "Hello World"
	if strings.HasPrefix(str, "Hello") {
		fmt.Println("Yes1")
	} else if strings.HasSuffix(str, "World") {
		fmt.Println("Yes2")
	}
	//字符串包含关系

}
