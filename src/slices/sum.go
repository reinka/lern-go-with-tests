package slices

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAllTails(numbersSlice ...[]int) []int {
	results := make([]int, len(numbersSlice))
	for idx, nums := range numbersSlice {
		sum := 0
		if len(nums) > 0 {
			sum = Sum(nums[1:])
		}
		results[idx] = sum
	}

	return results
}
