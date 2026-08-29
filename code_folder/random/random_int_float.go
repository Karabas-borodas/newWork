package main

import (
	// "crypto/rand"
	"fmt"
	// "math/big"
	"math/rand/v2"
)

func main() {
	//:NOTE:cripto random
	// n, err := rand.Int(rand.Reader, big.NewInt(100))
	// if err != nil {
	// 	fmt.Println("aaaaa random error", err)
	// }
	// fmt.Println("int n", n)

	//NOTE:standart randim
	rand.Seed(5)
	fmt.Println("int", rand.Int())
	fmt.Println("int", rand.Int())
	fmt.Println("IntN", rand.IntN(10))
	fmt.Println("float64", rand.Float64())
	fmt.Println("Nfloat64", rand.NormFloat64())
}
