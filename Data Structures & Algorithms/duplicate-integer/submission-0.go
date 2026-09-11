func hasDuplicate(nums []int) bool {
    m := make(map[int]int);

    for _, val := range nums {
        if(m[val] > 0){
            return true;
        }

        m[val]++;
    }

    return false; 
}
