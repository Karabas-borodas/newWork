package main

import (
	"fmt"
)

func main() {

	fmt.Println("-----1----")
	var Id int = 1
	var name string = "Aloha"
	var age int = 10
	user := NewUser(Id, name, age)
	fmt.Println(user)
	fmt.Println("-----2----")
	var x = 45
	var y = 67
	point := NewPoint(x, y)
	fmt.Printf("point %v", point)
	fmt.Println("-----3----")
	var r Rectangle
	r.Heigth = 4
	r.Width = 5
	fmt.Println(Area(r))
}
