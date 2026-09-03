package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value struct{}
	var in interface{}

	fmt.Println(unsafe.Sizeof(value))
	fmt.Println(unsafe.Sizeof(in))
}
