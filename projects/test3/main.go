package main

import "fmt"

func main() {
	var a = '测'
	fmt.Printf("%T %[1]d %[1]c\n", a)
	//a = 'asd' // 这么写不对 因为字符类型只能一个字符 这样叫字符串
	a = 'c' // 这样写对吗 a只需要2个字节，需要int32表示吗? 对因为是重新赋值，不是重新定义，所以类型不会修改
	fmt.Printf("%T %[1]d %[1]c\n", a)
	a = 99 // 可以重新赋值99 因为99也是int形
	fmt.Printf("%T %[1]d %[1]c\n", a)
	a = 0x63 // 整数表示法
	fmt.Printf("%T %[1]d %[1]c\n", a)
	a = '\x63' // 16进制表示法
	fmt.Printf("%T %[1]d %[1]c\n", a)

}

// 在计算机中一切数据都是 数字，我们赋予数字不同的类型，他就会显示不同的效果，本来数字99赋予他字符型他就是C
// 可以认为在Go语言中没有字符型，字符型本是上rune ，rune = int32

/*
 var a int32 = 123    右边如果是untyped constant，会根据左边指定类型转换，这种转换是隐式的
 var a int32 = int32 (123)  正确，指定了类型，也赋值。不会采用类型推到TODO，隐式类型转换  语法糖 这两种写法一样，上面那种用了隐式类型转换
*/
