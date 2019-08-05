package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	m := make(map[int]int, n) //hashtable
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		m[a] = a
	}

	countPairs := 0
	for key, _ := range m {
		diff := key - k
		if diff <= 0 {
			continue
		}

		if _, ok := m[diff]; ok {
			countPairs++
		}
	}

	fmt.Print(countPairs)
}
