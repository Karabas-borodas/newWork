package main

import (
	"fmt"
)

func main() {
	slice := []float32{2.14, 15, 33.2}
	var summ float32
	for _, v := range slice {
		summ += v
	}
	fmt.Println(summ)
}
