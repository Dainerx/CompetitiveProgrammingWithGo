package main

import (
	"fmt"
	"sort"
)

func search(nums []int, target int) int {
	sep := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			sep = i + 1
			break
		}
	}
	if sep != -1 {
		left, right := nums[0:sep], nums[sep:]
		sl, sr := sort.SearchInts(left, target), sort.SearchInts(right, target)
		if sl != len(left) {
			if left[sl] == target {
				return sl
			}
		}
		if sr != len(right) {
			if right[sr] == target {
				return sr + len(left)
			}
		}
	} else {
		sm := sort.SearchInts(nums, target)
		if sm != len(nums) {
			if nums[sm] == target {
				return sm
			}
		}
	}
	return -1
}

func main() {
	a := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(a, 8))
}
