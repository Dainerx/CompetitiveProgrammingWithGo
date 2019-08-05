package main

//if this gives TLE it is due to how slow fmt.Scanf
import "fmt"

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for i := 0; i < T; i++ {
		var n int
		fmt.Scanf("%d", &n)
		numbers := make([]int64, n)
		acc := make([]int64, n)
		fmt.Scanf("%d", &acc[0])
		numbers[0] = acc[0]
		for j := 1; j < n; j++ {
			var a int64
			fmt.Scanf("%d", &a)
			numbers[j] = a
			acc[j] = a + acc[j-1]
		}
		var current int64 = 0
		found := false
		for j := 0; j < n; j++ {
			if right := acc[len(acc)-1] - acc[0]; j == 0 && current == right {
				fmt.Println("YES")
				found = true
				break
			}
			if j != 0 {
				current = acc[j]
				left, right := acc[j-1], acc[len(acc)-1]-current
				if left == right {
					fmt.Println("YES")
					found = true
					break
				}
			}
		}
		if found == false {
			fmt.Println("NO")
		}
	}
}
