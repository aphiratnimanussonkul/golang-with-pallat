package foobar

import (
	"strconv"
)

func Say(num int) string {

	if num == 6 {
		return "Foo"
	}
	if num == 5 {
		return "Bar"
	}
	if num == 3 {
		return "Foo"
	}
	return strconv.Itoa(num)
}

func SayAny(num interface{}) string {
	var n int
	if s, ok := num.(string); ok {
		n, _ = strconv.Atoi(s)
	}
	if i, ok := num.(int); ok {
		n = i
	}
	return Say(n)
}

func SayAny2(num interface{}) string {
	var n int
	switch v := num.(type) {
	case int:
		n = v
	case string:
		n, _ = strconv.Atoi(v)
	default:
	}

	return Say(n)
}
