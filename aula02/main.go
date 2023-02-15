package main

import "fmt"

func main() {
	fmt.Println(calcMin(1, 2, 3, 4, 5))
}
func calcMin(n ...int) int {
	min := n[0]

	for _, num := range n {
		if num < min {
			min = num
		}
	}
	return min
}
