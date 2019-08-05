package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	points := []int{
		a, b, c,
	}
	sort.Ints(points)
	answer := (points[1] - points[0]) + (points[2] - points[1])
	fmt.Println(answer)
}
