func twoSum(nums []int, target int) []int {
    var map1 = make(map[int]int);
	for i:=0;i < len(nums); i++ {
		val, ok := map1[nums[i]]
		if ok {
			return []int{val, i};
		} else {
			map1[target - nums[i]] = i;
		}
	}

	return nil;
}
