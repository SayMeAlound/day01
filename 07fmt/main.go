package main

import "fmt"

//fmt 占位符
func main() {

	var n = 100
	fmt.Printf("%T\n", n) //查看变量类型
	fmt.Printf("%v\n", n) //查看变量值
	fmt.Printf("%b\n", n) //二进制
	fmt.Printf("%o\n", n) //八进制
	fmt.Printf("%d\n", n) //十进制
	fmt.Printf("%x\n", n) //16进制

	var s = "go go go  masiwei"
	fmt.Printf("字符串: %s\n", s)  //打印字符串
	fmt.Printf("字符串: %v\n", s)  //查看变量值(v  value  通用)
	fmt.Printf("字符串: %#v\n", s) //加描述符

}
