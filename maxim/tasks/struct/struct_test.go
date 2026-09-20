package main

import (
	"fmt"
	"testing"
)

func Test4_1(t *testing.T) {
	var point1 Point = Point{X: 0,
		Y: 0}
	b := IsOrigin(point1)
	if b == false {
		t.Errorf("%d!=0,%d!=o", point1.X, point1.Y)
	}
}
func Test4_2(t *testing.T) {
	var point1 Point = Point{X: 3,
		Y: 2}
	b := IsOrigin(point1)
	if b == true {
		t.Errorf("%d==0,%d==o", point1.X, point1.Y)
	}
}
func Test6_1(t *testing.T) {
	var user User = User{ID: 15,
		Name: "ALah",
		Age:  3,
	}

	user2 := WithEmail(user, "5553535")
	if user.ID != user2.ID || user.Name != user2.Name || user.Age != user2.Age || user2.Email != "5553535" {
		t.Errorf("%d==0,%d==o", user.Age, user2.Age)
	}
}

func Test6_2(t *testing.T) {
	user := User{ID: 15, Name: "ALah", Age: 3}

	user2 := WithEmail(user, "5553535")
	user2.ID = 10

	if user.ID != 15 {
		t.Errorf("original mutated: user.ID = %d, want 15", user.ID)
	}
	if user2.ID != 10 {
		t.Errorf("copy not modified: user2.ID = %d, want 10", user2.ID)
	}
	if user2.Email != "5553535" {
		t.Errorf("user2.Email = %q, want %q", user2.Email, "5553535")
	}
	if user.Email != "" {
		t.Errorf("original Email mutated: %q", user.Email)
		fmt.Printf("user  = %+v\n", user)
		// fmt.Printf("user2 = %+v\n", user2)
	}

}
