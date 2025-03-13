package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0
	}
	lastWord := words[len(words)-1]
	return len(lastWord)
}

func main() {
	input := "Hello world from Go"
	result := lengthOfLastWord(input)
	fmt.Println("Length of the longest word:", result)
}
