package main

import "math"

func searchInsert(nums []int, target int) int {
	r := len(nums) - 1
	l := 0
	ans := -1
	for r >= l {
		mid := math.Floor(float64((r + l) / 2))
		if nums[int(mid)] == target {
			ans = int(mid)
			r = int(mid) - 1
		} else if nums[int(mid)] > target {
			r = int(mid) - 1
		} else {
			l = int(mid) + 1
		}
	}
	if ans != -1 {
		return ans
	}
	return l
}

func main() {

}
