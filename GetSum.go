// First project main.go
package main

import (
	"fmt"
)

func swap(a int, b int) (int, int) {
	return b, a
}

func add(x int) int {
	x += 1
	return x
}

func getsum(num []int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}

func addc(x *int) *int {
	*x += 1
	return x
}

func main() {
	num := []int{1, 2, 3, 4, 5}
	fmt.Print(getsum(num))
}
