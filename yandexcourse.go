package main

import (
	"fmt"
	"strings"
)

func Mul(a interface{}, b int) interface{} {
	switch a2 := a.(type) {
	case int:
		return a2 * b
	case string:
		return strings.Repeat(a2, b)
	case fmt.Stringer:
		return strings.Repeat(a2.String(), b)
	default:
		panic("unrecognized type; only string, Stringer and int are accepted")
	}
}

func main() {
	fmt.Println(Mul(1, 5))
	fmt.Println(Mul("12", 5))
	fmt.Println(Mul(nil, 5))
}
