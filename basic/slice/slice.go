package slice

import (
	"fmt"
)

func MainSlice() {
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a := array[0:5:6]
	fmt.Println(a[:], cap(a), len((a)))

	xx(1, 2, 3, 4, 5)
	yy([]int{1, 2, 3, 4, 5})
}

func Couple(str string) []string {
	var r []string
	s := []rune(str)
	for s = append(s, []rune("*")...); len(s) > 1; s = s[2:] {
		r = append(r, string(s[:2]))
	}
	return r
}

func xx(a ...int) {}
func yy(a []int)  {}

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8
	for i := 0; i < dx; i++ {
		var row []uint8
		for j := 0; j < dy; j++ {
			row = append(row, uint8(i*j/2))
		}
		pic = append(pic, row)
	}
	return pic
}
