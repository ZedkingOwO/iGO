package main

import "fmt"

func main() {

	const (
		a = 100
		b = 200
	)

	const (
		q = iota + 1 // 批量定义下面可以继承省略
		w
		e
		_ // 空白标识 跳过了3
		d
	)
	fmt.Println(q, w, e, d)
}
