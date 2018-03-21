package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("please provide two strings")
		return
	}

	str1, str2 := os.Args[1], os.Args[2]

	if str1 == str2 {
		fmt.Printf("%s and %s are the same string\n", str1, str2)
		return
	}

	if isAnagram(os.Args[1], os.Args[2]) {
		fmt.Printf("%s is an anagram of %s\n", str1, str2)
		return
	}

	fmt.Printf("%s is not an anagram of %s\n", str1, str2)
}

func isAnagram(s1, s2 string) bool {
	r1 := []rune(s1)
	r2 := []rune(s2)
	if len(r1) != len(r2) {
		return false
	}

	runeMap1 := make(map[rune]int)
	runeMap2 := make(map[rune]int)

	countRunes(r1, runeMap1)
	countRunes(r2, runeMap2)

	return compareRunes(runeMap1, runeMap2)
}

func countRunes(runes []rune, rMap map[rune]int) {
	for _, r := range runes {
		rMap[r]++
	}
}

func compareRunes(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for r, c := range m1 {
		if c != m2[r] {
			return false
		}
	}

	return true
}
