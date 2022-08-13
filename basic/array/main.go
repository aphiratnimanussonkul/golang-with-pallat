package array

import "fmt"

func MainArray() {
	matrix := [4][4]int{
		{1, 3, 4, 2},
		{3, 4, 2, 1},
		{4, 2, 1, 3},
		{2, 1, 3, 4},
	}
	fmt.Println(multiply(matrix, 4))
}

func multiply(matrix [4][4]int, multiply int) [4][4]int {
	for row := range matrix {
		for col := range matrix[row] {
			matrix[row][col] *= multiply
		}

	}
	return matrix
}
