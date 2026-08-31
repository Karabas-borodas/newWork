package main

import (
	"fmt"
)

// NOTE:1
func CountInts(nums []int) map[int]int {
	ma := make(map[int]int)
	for _, v := range nums {
		ma[v] += 1
	}
	return ma
}
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 2, 3, 1, 2}
	fmt.Println(slice)
	fmt.Println("#1")
	fmt.Println(CountInts(slice))
}
