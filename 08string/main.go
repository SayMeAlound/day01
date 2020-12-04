package main

import (
	"fmt"
	"strings"
)

//字符串

func main() {
	// 本来具有特殊含义的,我应该告诉程序我写反斜线就是一个单纯的反斜线
	path := "\"C:\\Users\\zhaomo\\Desktop\\PyDemo\""
	fmt.Println(path)

	s := "I'm ok"
	fmt.Print(s)

	//多行的字符串
	s2 := `
		为什么
		我是你的
		星星`
	fmt.Println(s2)

	//字符串相关操作
	fmt.Println(len(s2))

	//字符串拼接
	chinese := "理性"
	math := "感性"
	man := math + chinese
	fmt.Println(chinese + math)
	fmt.Println(man)

	//拼接后打印
	fmt.Printf("%s%s", chinese, math)
	fmt.Println("------------")

	//拼接后返回一个字符串变量
	ss := fmt.Sprintf("%s%s", chinese, math)
	fmt.Println(ss)
	//分隔
	ret := strings.Split(path, "\\")
	fmt.Println(ret)
	//判断包含
	fmt.Println(strings.Contains(chinese, "理性"))

	//前缀
	fmt.Println(strings.HasPrefix(chinese, ";理"))
	//后缀
	fmt.Println(strings.HasSuffix(chinese, "性"))

	//字串出现的位置
	s4 := "In Your Eyes"
	fmt.Println(s4)
	fmt.Println(strings.Index(s4, "E"))

	//拼接
	fmt.Println(strings.Join(ret, "+"))

}
