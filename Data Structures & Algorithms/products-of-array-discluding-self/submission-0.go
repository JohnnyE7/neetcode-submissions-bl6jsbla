func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	res, pref, post := make([]int, n), make([]int, n), make([]int, n)
	
	pref[0] = 1
	for i := 1; i < n; i++ {
		pref[i] = pref[i-1] * nums[i-1]
	}
	
	post[n-1] = 1
	for i := n-2; i >= 0; i-- {
		post[i] = post[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		res[i] = pref[i] * post[i]
	}

	return res
}

