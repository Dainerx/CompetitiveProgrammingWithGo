package main

import (
	"fmt"
	"sort"
)

//if this gives TLE it is due to how slow fmt.Scanf

type numberIndex struct {
	number int
	index  int
}

func getMaxAndCut(cons []numberIndex) []numberIndex {
	tmp := make([]numberIndex, len(cons))
	copy(tmp, cons)
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].number < tmp[j].number
	})
	indexToCutFrom := tmp[len(tmp)-1].index
	return cons[:indexToCutFrom]
}

func main() {
	var g int8
	fmt.Scanf("%d", &g)
	for g > 0 {
		g--
		var n int
		fmt.Scanf("%d", &n)
		numbers := make([]numberIndex, n)
		for i := 0; i < n; i++ {
			var a int
			fmt.Scanf("%d", &a)
			ni := numberIndex{a, i}
			numbers[i] = ni
		}
		games := 0
		for len(numbers) > 0 {
			numbers = getMaxAndCut(numbers)
			games++
		}
		if games%2 == 1 {
			fmt.Println("BOB")
		} else {
			fmt.Println("ANDY")
		}
	}
}
