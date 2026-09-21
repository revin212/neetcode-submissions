func topKFrequent(nums []int, k int) []int {
	// Step 1: count frequency of each number
	freqMap := make(map[int]int)
	for _, n := range nums {
		freqMap[n]++
	}

	// Step 2: create buckets indexed by frequency
	// bucket[i] holds all numbers that appear exactly i times
	// max possible frequency is len(nums)
	buckets := make([][]int, len(nums)+1)
	for num, freq := range freqMap {
		buckets[freq] = append(buckets[freq], num)
	}

	// Step 3: iterate buckets from highest frequency to lowest
	result := make([]int, 0, k)
	for freq := len(buckets) - 1; freq >= 0 && len(result) < k; freq-- {
		for _, num := range buckets[freq] {
			result = append(result, num)
			if len(result) == k {
				break
			}
		}
	}

	return result
}