package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println("int", rand.IntN(1000))
	}
	fmt.Println("IntN", rand.IntN(10))
	fmt.Println("float64", rand.Float64())
	fmt.Println("Nfloat64", rand.NormFloat64())
}
