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

	//声明变量同时赋值
	var s1 string = "whb"
	fmt.Println(s1)
	//类型推导(根据值判断该变量是什么类型)
	var s2 = "20"
	fmt.Println(s2)
	//简短变量声明  只能在函数里面用
	s3 := "哈哈哈"
	fmt.Print(s3)
}
