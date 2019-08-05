package main

import "fmt"

func countInversions(arr []int) int {
	inversions := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				inversions++
			}
		}
	}
	return inversions
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			var a int
			fmt.Scanf("%d", &a)
			arr[j] = a
		}
		if countInversions(arr)%2 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
