package main

import (
	"fmt"
	"strings"
)

func countVowels(text string) int {
	var count int
	for _, char := range text {
		switch char {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			count++
		}
	}
	return count
}

func reverseString(text string) string {
	runeValue := []rune(text)
	for i, j := 0, len(runeValue)-1; i < j; i, j = i+1, j-1 {
		runeValue[i], runeValue[j] = runeValue[j], runeValue[i]
	}
	return string(runeValue)
}

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func main() {
	text := "Hello!"
	fmt.Println(text)
	fmt.Println(len(text))
	fmt.Println(len([]rune(text)))

	message := "Go😀Lang"
	for index, value := range message {
		fmt.Println(index, value, string(value))
	}

	fmt.Println(countVowels("I am genius"))

	fmt.Println(reverseString("SIRENHEAD 💀"))

	fmt.Println(countWords("Programming in go is really fun and I'm a geek Gopher! 🙃 "))
}
