package main

import (
	"fmt"
	"sort"
)
/*
//this is much faster given the duplicate numbers
func search(nums []int, target int) bool {
	nmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		nmap[nums[i]] = -1
	}
	if _, ok := nmap[target]; ok {
		return true
	}
	return false
}
*/
func search(nums []int, target int) bool {
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
				return true
			}
		}
		if sr != len(right) {
			if right[sr] == target {
				return true
			}
		}
	} else {
		sm := sort.SearchInts(nums, target)
		if sm != len(nums) {
			if nums[sm] == target {
				return true
			}
		}
	}
	return false
}

func main() {
	a := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(a, 8))
}
