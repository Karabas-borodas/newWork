package main

import (
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
