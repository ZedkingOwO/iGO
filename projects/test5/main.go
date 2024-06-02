package main

import "fmt"

func main() {
	// make函数 构建切片、map、chan
	var c4 = make([]int, 1, 5)
	fmt.Printf("%d %d %d\n",c4,len(c4),cap(c4))

}
