package demo01

import "fmt"

func Swap(a, b int) {
	a, b = b, a
	fmt.Println("a:", a, "b:", b)
}

func SwapPoint(a, b *int) {
	*a, *b = *b, *a
}
