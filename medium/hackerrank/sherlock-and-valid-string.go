package main

import (
	"fmt"
	"math"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)
	occurence := make(map[rune]int)
	for _, c := range s {
		if _, ok := occurence[c]; ok {
			occurence[c] = occurence[c] + 1
		} else {
			occurence[c] = 1
		}
	}

	max, min, occMax, occMin := math.MinInt32, math.MaxInt32, 0, 0
	for _, val := range occurence {
		if val >= max {
			if val > max {
				occMax = 1
			} else {
				occMax++
			}
			max = val
		}
		if val <= min {
			if val < min {
				occMin = 1
			} else {
				occMin++
			}
			min = val
		}
	}

	if max == min {
		fmt.Println("YES")
	} else if max-min == 1 && (occMax == 1 || occMin == 1) {
		fmt.Println("YES")
	} else if ((occMax * max) == len(s)-1) && min == 1 && occMin == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
