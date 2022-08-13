package mapp

import "fmt"

func MainMapp() {
	m := map[string]int{
		"G": 71,
		"O": 79,
		"P": 80,
		"H": 72,
		"E": 69,
		"R": 82,
	}
	keys := []string{}
	vals := []int{}

	for k, v := range m {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	fmt.Println(keys)
	fmt.Println(vals)
}
