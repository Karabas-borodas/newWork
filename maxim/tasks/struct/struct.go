package main

import (
	"fmt"
	"math"
)

// NOTE:1
// Определите структуру User с полями ID int, Name string, Age int. Напишите функцию NewUser(id int,
// name string, age int) User, которая создает и возвращает пользователя.

type User struct {
	ID    int
	Name  string
	Age   int
	Email string
}

func NewUser(id int, name string, age int) *User {

	if age < 0 {
		fmt.Errorf("age cant be lou 0")
	}
	var user User
	user.ID = id
	user.Age = age
	user.Name = name
	return &user
}

// NOTE:2
// Определите структуру Point с координатами X int и Y int. Напишите функцию NewPoint(x, y int)
// Point.
type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	var p Point
	p.X = x
	p.Y = y
	return &p
}

// NOTE:3
// Определите структуру Rectangle с полями Width int и Height int. Напишите функцию Area(r
// Rectangle) int.
type Rectangle struct {
	Width  int
	Heigth int
}

func Area(x Rectangle) int {
	if x.Width < 0 || x.Heigth < 0 {
		return -1
	}
	if x.Heigth > 0 && x.Width > math.MaxInt/x.Heigth {

		return -1
	}
	ar := x.Heigth * x.Width
	return ar
}

// NOTE:4
// Для структуры Point напишите функцию IsOrigin(p Point) bool, которая возвращает true, если обе
// координаты равны нулю.
func IsOrigin(p Point) bool {
	if p.X == 0 && p.Y == 0 {
		return true
	}
	return false
}

// NOTE:6
// Для структуры User добавьте поле Email string. Напишите функцию WithEmail(u User, email string)
// User, которая возвращает измененную копию пользователя и не меняет исходный u.
func WithEmail(u User, email string) User {
	u.Email = email
	return u
}

// NOTE:7
type Pair struct {
	A int
	B int
}

func SwapPair(A int, B int) (int, int) {
	A, B = B, A
	return A, B
}
