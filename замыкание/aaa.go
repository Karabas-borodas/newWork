package main

import (
	"fmt"
	// "os"
)

// FIX: почему так происходит где заканчивается контекскт и начинается новый?
func main() {
	adder := func(variable int) func(int) int {
		return func(z int) int {
			return variable + z
		}
	}
	add := adder(10)
	add1 := adder(add(5))
	fmt.Println(add1(10))
}

// func main() {
//
// 	minus := func(variable int) func(int) int {
// 		return func(v int) int {
// 			return variable - v
// 		}
// 	}
// 	min := minus(6)
// 	fmt.Println(min(3))
// 	fmt.Println(min(3))
// }
