package main

import (
	"testing"
)

func Test18_0(t *testing.T) {
	str1 := "aA"
	str2 := "Aa"
	result := AreAnagrams(str1, str2)

	// Если "aA" и "Aa" НЕ анаграммы (регистрозависимые)
	if result == false {
		t.Errorf("expected false, got true for %s - %s", str1, str2)
	}

	// Если "aA" и "Aa" анаграммы (регистронезависимые)
	// if result == false {
	// 	t.Errorf("expected true, got false for %s - %s", str1, str2)
	// }
}
