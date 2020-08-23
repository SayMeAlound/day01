package main

import "fmt"

//声明变量
//var name string
//var age int
//var isOk bool

//批量声明变量
var (
	name string //""
	age  int    //0
	isOk bool   //false
)

func main() {
	name = "zhaomo"
	age = 18
	isOk = true

	fmt.Print(isOk)             //在终端中打印出内容
	fmt.Println(age)            //打印完指定的内容之后会在后面加一个换行符
	fmt.Printf("name:%s", name) //%s是占位符 后面的name替换占位符
}
