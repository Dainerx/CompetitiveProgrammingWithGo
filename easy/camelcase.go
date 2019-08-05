package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)
	words := 1
	countWords := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			words++
			return r
		default:
			return r
		}
	}

	strings.Map(countWords, s)
	fmt.Println(words)

}
