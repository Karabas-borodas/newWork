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

	//NOTE:я сделяль
	// l := len(r)
	// fmt.Println(l)
	// if l%2 != 0 {
	// 	l -= 1
	// 	for i := 0; i < (len(r))/2; i++ {
	// 		if r[i] != r[l-i] {
	// 			return false
	// 		}
	// 	}
	// } else {
	//
	// 	for i := 0; i < (len(r))/2; i++ {
	// 		if r[i] != r[l-1-i] {
	// 			return false
	// 		}
	// 	}
	// }

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}

	return true
}

// NOTE:10
func NormalizeSpaces(s string) string {
	count := 0
	r := []rune(s)
	for i := 0; i < len(r)-1; i++ {
		if s[i] == ' ' {
			count += 1
		} else {
			break
		}

	}
	r = r[count:]
	for i := len(r) - 1; i > 0; i-- {
		if r[i] == ' ' {
			r = r[:i]
		} else {
			break
		}

	}
	return string(r)
}

// NOTE:12
func IndexOf(s, sub string) int {
	if len(s) == 0 {
		return -1

	}
	if len(sub) > len(s) {
		return -1
	}
	subIndex := 0
	for i := 0; i < len(s); i++ {
		if s[i] == sub[subIndex] && subIndex < len(sub)-1 {
			subIndex += 1
		} else if subIndex == len(sub)-1 {
			return i - subIndex
		} else {
			subIndex = 0
		}

	}
	return -1
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
	fmt.Println("#10")
	fmt.Println(NormalizeSpaces(" ШаЛаШ"))
	fmt.Println(NormalizeSpaces("   ШаЛаШ"))
	fmt.Println(NormalizeSpaces("аЛ   аШ  "))
	fmt.Println(NormalizeSpaces("     "))
	fmt.Println("#12")
	fmt.Println(IndexOf(string, "дела"))
}
