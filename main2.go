// First project main.go
package main

import (
	"fmt"
)

var f complex64 = 2 + 2i

func swap(a int, b int) (int, int) {
	return b, a
}

func add(x int) int {
	x += 1
	return x
}

func addc(x *int) *int {
	*x += 1
	return x
}

func main() {
	a := 1
	//b := 2
	addc(&a)
	fmt.Print(a)
}
