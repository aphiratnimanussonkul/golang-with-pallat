package main

import (
	"fmt"

	"github.com/aphiratnimanussonkul/golang-with-pallat/basic/interfacee"
)

var name string = "Aphirat"

func main() {
	interfacee.MainAnimalColor()
}

func isPalindrom(a, b, c, d, e int) bool {
	return a == e && b == d
}

func power(b, x int) int {
	result := 1
	for i := 0; i < x; i++ {
		result = result * b
	}
	return result
}

func powerRecursive(b, x int) int {
	if x > 0 {
		return b * powerRecursive(b, x-1)
	}
	return 1
}

func mainBasic() {
	println("Hello, world!")
	fmt.Println("My area", squareArea((4)))

	if ok := isCorrct(); ok {
		fmt.Println("It's correct")
	}
}

func squareArea(a float64) float64 {
	return a * a
}

func isCorrct() bool {
	return true
}
