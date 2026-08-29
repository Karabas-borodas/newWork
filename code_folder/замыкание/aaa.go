package main

import (
	"fmt"
	// "os"
)

// FIX: почему так происходит где заканчивается контекскт и начинается новый?
func main() {
	adder := func(variable int) func(int) int {
		return func(z int) int {
			variable += z
			return variable + z
		}
	}
	//NOTE: когза я создаю функцию он захыватывает переменную в первую функцию
	//NOTE:почему в первую? почему не во вторую?

	add := adder(10)
	//NOTE:когда я еще раз ее использую то почему переменная попадает во втору Z а не заменяет VARIABLE?

	//NOTE: почему при повтормно вызове add() он уже не захыватывает переменную и выводит адресс памяти?
	fmt.Printf("%d\n", add(5))
	fmt.Printf("%d\n", add(5))
	fmt.Printf("%d\n", add(5))
	fmt.Printf("%d\n", add(5))
	// add1 := adder(add(5))
	// fmt.Println(add1(10))
}
