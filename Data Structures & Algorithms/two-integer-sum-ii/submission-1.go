func twoSum(numbers []int, target int) []int {
	leftPointer := 0
	rightPointer := len(numbers)-1
	temp := 0

	for leftPointer < rightPointer {
		temp = numbers[leftPointer] + numbers[rightPointer] 
		if temp == target {
			return []int {leftPointer+1, rightPointer+1}
		} else if temp > target {
			rightPointer--
		} else {
			leftPointer++
		}
	}

	return []int {}
}
