package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		// fmt.Print(i)
		// fmt.Print(i)
		// fmt.Print(i)
		// fmt.Print(i)
		// fmt.Print(i)
		// fmt.Print(i)
		// fmt.Print(i)
		// fmt.Print(i)
		// fmt.Print(i)

		for j := 1; j <= i; j++ {
			fmt.Printf("%d%c%d%c%d\t", i, '*', j, '=', i*j)

		}
		fmt.Println()
	}
}
