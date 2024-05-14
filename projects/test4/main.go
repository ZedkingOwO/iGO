package main

import "fmt"

func main() {
	var s = 27979 // 6 x 16^3 + 13x16^2 + 4x16^1 + 11 =
	fmt.Printf("%+v %[1]d %[1]b %[1]o  %[1]x %[1]X %[1]c %[1]U\n", s)
	t := fmt.Sprintf("%+v %[1]d %[1]b %[1]o  %[1]x %[1]X %[1]c %[1]U\n", s)
	//fmt.print(s) 缺省
	fmt.Println(t)
	// 生产字符串，然后输出到标准输出
}
