import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break // smallest number is positive, no sum can reach 0
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate first number
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++ // skip duplicate second number
				}
				for l < r && nums[r] == nums[r+1] {
					r-- // skip duplicate third number
				}
			}
		}
	}

	return result
}