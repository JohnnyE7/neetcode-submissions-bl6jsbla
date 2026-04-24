func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	res := make([]int, n)
	
	pref := 1
	for i := 0; i < n; i++ {
		res[i] = pref
		pref *= nums[i]
	}
	
	post := 1
	for i := n-1; i >= 0; i-- {
		res[i] *= post
		post *= nums[i]
	}

	return res
}

