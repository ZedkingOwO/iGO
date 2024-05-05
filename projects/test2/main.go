package main

import "fmt"

func main() {
	//	var a = 100
	//	var b int64 = 200
	//	var c rune = 300 // type rune = init32  类型别名rune  他是4tytes或32bits int

	/* 这两个的意思完全不同
	type B = int32    q 指代 int32   q可以和int32类的值互操作      类型别名
	type x int32	  我创建了一个型类型X 他的类型是int32新的类型	重新定义 虽然你的本质也是int32 但是你是新的类型
	*/
	type x int32
	type B = int32
	var G B = 100
	var m x = 111
	//	fmt.Println(G + m)
	fmt.Println(G + int32(m))
}
