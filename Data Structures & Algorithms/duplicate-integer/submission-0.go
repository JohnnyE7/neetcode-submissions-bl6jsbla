func hasDuplicate(nums []int) bool {
    set := make(map[int]struct{}, len(nums)- 1)
	
	for _, num := range nums {
		set[num] = struct{}{}
	}

	if len(nums) != len(set) { return true }
	return false 
}
