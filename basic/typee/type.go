package typee

import (
	"fmt"
	"strconv"
	"strings"
)

type Int int

type String string

func (s String) toInt() int {
	i, _ := strconv.Atoi(string(s))
	return i
}

func (s *String) toUppper() {
	*s = String(strings.ToUpper(string(*s)))
}

func MainType() {
	var numInt Int = 10
	fmt.Printf("%T %q\n", numInt, numInt)
	fmt.Println(numInt.String())

	var s String = "Aphirat"
	s.toUppper()
	fmt.Println(s)

	var sInt String = "20"
	fmt.Println(sInt.toInt())

}

func (i Int) String() string {
	return strconv.Itoa(int(i))
}
