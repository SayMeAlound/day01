package main

import "fmt"

//常量
const pi = 3.1415926

//批量声明常量
const (
	STATUSOK = 200
	NOTFOUND = 404
)

//批量声明常量时  如果每一行声明后 没有赋值  默认和上一行值一致
const (
	n1 = 100
	n2
	n3
)

//iota
const (
	a1 = iota //0
	a2        //1
	a3        //2
)
const (
	b1 = iota //0
	b2 = iota //1
	_  = iota //2
	b3        //3
)

//const中每新增一行常量声明将使 iota计数一次(iota可理解为const语句块中的行索引)
const (
	c1 = iota //0
	c2 = 100  //100
	c3
	c4
)

//多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 //d1: 1  d2:2
	d3, d4 = iota + 1, iota + 2 //d3: 2  d4:3
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) //1 左移10位  等于  100000000000   等于10进制的1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	//fmt.Println(a1)
	//fmt.Println(a2)
	//fmt.Println(a3)

	//fmt.Println(b1)
	//fmt.Println(b2)
	//fmt.Println(b3)
	//
	//fmt.Println(c1)
	//fmt.Println(c2)
	//fmt.Println(c3)
	//fmt.Println(c4)

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)

}
