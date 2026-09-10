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
		} else if subIndex == len(sub)-1 && s[i] == sub[subIndex] {
			return i - subIndex
		} else {
			subIndex = 0
		}

	}
	return -1
}

// TODO:: после созвона 03.09.26
// NOTE:14
// Подсчитайте все вхождения подстроки, включая пересекающиеся. Например, в "aaaa"
// строка "aa" встречается три раза.
func CountOverlapping(s, sub string) (int, error) {
	if len(sub) > len(s) {
		return 0, fmt.Errorf(" sub struct longer first")
	}
	if len(sub) == 0 || len(s) == 0 {
		return 0, fmt.Errorf("empty strung")
	}
	runS := []rune(s)
	runeSub := []rune(sub)
	count := 0
	for i := 0; i < len(runS); i++ {
		if runS[i] == runeSub[0] {
			var flag bool = true
			for j := 0; j < len(runeSub); j++ {
				if i+j >= len(runS) {
					flag = false
					break
				}
				if runS[i+j] != runeSub[j] {
					flag = false
					break
				}
			}
			if flag == true {
				count++
			}

		}
	}
	return count, nil
}

// NOTE:16
// Найдите самый длинный общий префикс всех строк в слайсе. Для пустого слайса
// верните пустую строку.
func LongestCommonPrefix(items []string) string {
	if len(items) == 1 {
		return items[0]
	}
	if len(items) == 0 {
		return ""
	}
	var str = items[0]
	for _, c := range items {
		if len(str) > len(c) {
			str = c
		}
	}
	count := 0
	for i := 1; i < len(items); i++ {
		for j := 0; j < len(str); j++ {
			if str[j] != items[i][j] {
				break

			} else {
				count++

			}
		}
		str = str[:count]
		count = 0
		// fmt.Println(len(str))
		// str = count
	}

	return str
}

func main() {
	var stringers = "Странные дела! Awada Kedawra \n OP musorok ne shey mne SroKK"
	fmt.Println(stringers)
	fmt.Println("#1")
	fmt.Println(StringSize(stringers))
	fmt.Println("#3")
	fmt.Println(FirstRune(stringers))
	fmt.Println("#8")
	fmt.Println(IsPalindrome("ШаЛаШ"))
	fmt.Println(IsPalindrome("аЛаШ"))
	fmt.Println("#10")
	fmt.Println(NormalizeSpaces(" ШаЛаШ"))
	fmt.Println(NormalizeSpaces("   ШаЛаШ"))
	fmt.Println(NormalizeSpaces("аЛ   аШ  "))
	fmt.Println(NormalizeSpaces("     "))
	fmt.Println(NormalizeSpaces(" a b "))
	fmt.Println("#12")
	fmt.Println(IndexOf(stringers, "дела"))
	fmt.Println(IndexOf("ab", "x"))
	fmt.Println("#14")
	fmt.Println(CountOverlapping("aaaa", "aa"))
	fmt.Println(CountOverlapping("aaaa", "aa"))
	fmt.Println(CountOverlapping("aaaa", ""))
	fmt.Println(CountOverlapping("aa  aa", "aa"))
	fmt.Println("#16")
	fmt.Println(LongestCommonPrefix([]string{"aaa", "aasss", "aaab"}))
	fmt.Println(LongestCommonPrefix([]string{"", "aasss", "aaab"}))
	fmt.Println(LongestCommonPrefix([]string{"abc", "aasss", "aaab"}))
}
