func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	n := uint(len(nums))
	for i := 0; i < 1<<n; i++ {
		a := []int{}
		for j := 0; j < int(n); j++ {
			if (i & (1 << uint(j))) > 0 {
				a = append(a, nums[j])
			}
		}
		result = append(result, a)
	}
	return result
}