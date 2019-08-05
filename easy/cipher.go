package main

import (
	"fmt"
)

func rotate(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

func main() {
	var n, k int
	var s string
	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &s)
	fmt.Scanf("%d", &k)

	var result []rune
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			result = append(result, rotate(r, 'A', k))
		} else if r >= 'a' && r <= 'z' {
			result = append(result, rotate(r, 'a', k))
		} else {
			result = append(result, r)
		}
	}

	fmt.Print(string(result))
}
