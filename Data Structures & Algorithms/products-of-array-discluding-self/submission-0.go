func productExceptSelf(nums []int) []int {
	arrLen := len(nums)

	result := make([]int, arrLen)
	prefixArr := make([]int, arrLen)
	suffixArr := make([]int, arrLen)

	//make prefix array :
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			prefixArr[0] = nums[0]
			prefixArr[1] = nums[1] * nums[0]

			suffixArr[arrLen-1] = nums[arrLen-1]
			suffixArr[arrLen-2] = nums[arrLen-2] * nums[arrLen-1]
		} else {
			prefixArr[i] = nums[i] * prefixArr[i-1]
			suffixArr[arrLen-1-i] = nums[arrLen-1-i] * suffixArr[arrLen-i]
		}
	}
	// fmt.Println(prefixArr)
	// fmt.Println(suffixArr)

	result[0] = suffixArr[1]
	for i := 1; i < len(nums)-1; i++ {
		result[i] = prefixArr[i-1] * suffixArr[i+1]
	}
	result[0] = suffixArr[1]
	result[arrLen-1] = prefixArr[arrLen-2]

	//make result array :
	return result
}