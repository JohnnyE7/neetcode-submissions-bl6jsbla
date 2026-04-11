func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	buckets := make([][]int, len(nums)+1)

	for _, num := range nums {
		freq[num]++
	}

	for key, val := range freq {
		buckets[val] = append(buckets[val], key)
	}

	res := make([]int, 0, k)
	for i := len(buckets)-1; i >= 0; i-- {
		for _, elem := range buckets[i] {
			res = append(res, elem)
			if len(res) == k {
				return res
			}
		}
	}
	
	return res
}
