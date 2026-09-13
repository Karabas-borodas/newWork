package main

import (
	"fmt"
)

func main() {

	var Id int = 1
	var name string = "Aloha"
	var age int = 10
	user := NewUser(Id, name, age)
	fmt.Println(user)
}
