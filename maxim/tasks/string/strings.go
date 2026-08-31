package main

import (
	"fmt"
	// "strconv"
)

// NOTE:1
func StringSize(s string) (bytes int, runes int) {
	bite := len(s)
	countRunes := []rune(s)
	return bite, len(countRunes)
}

// NOTE:3
func FirstRune(s string) (rune, bool) {

	if len(s) == 0 {
		var r rune
		return r, false
	} else {
		r := []rune(s)
		return rune(r[0]), true
	}
}

// NOTE:8
func IsPalindrome(s string) bool {

	r := []rune(s)

	l := len(r)
	fmt.Println(l)
	if l%2 != 0 {
		l -= 1
	}
	for i := 0; i < len(r)-1/2; i++ {
		if r[i] != r[l] {
			return false
		}
	}
	return true
}
func main() {
	var string = "Странные дела! Awada Kedawra \n OP musorok ne shey mne SroKK"
	fmt.Println(string)
	fmt.Println("#1")
	fmt.Println(StringSize(string))
	fmt.Println("#3")
	fmt.Println(FirstRune(string))
	fmt.Println("#8")
	fmt.Println(IsPalindrome("ШаЛаШ"))
	fmt.Println(IsPalindrome("аЛаШ"))
}
