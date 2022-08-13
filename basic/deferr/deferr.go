package deferr

import "fmt"

func MainDefer() {
	printInt(9)
}

func printInt(n int) {
	defer fmt.Println(n)
	fmt.Println(n + 1)
}

func MainRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("ok")
		}
	}()
	var s []int
	s[0] = 1
	fmt.Println(s)
}
