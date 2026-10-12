func maxArea(heights []int) int {
	result := 0
	l := 0
	r := len(heights)-1

	for l < r {
		area := (r-l) * min(heights[l],heights[r])
		if area > result {
			result = area
		}

		if heights[l] < heights[r] {
			l++
		} else if heights[l] > heights[r] {
			r--
		} else if heights[l] == heights[r] {
			l++
		}
	}

	return result
}
