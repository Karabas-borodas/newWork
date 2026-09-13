package main

import (
	"fmt"
)

// NOTE:1
// Определите структуру User с полями ID int, Name string, Age int. Напишите функцию NewUser(id int,
// name string, age int) User, которая создает и возвращает пользователя.

type User struct {
	ID   int
	Name string
	Age  int
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
