func longestConsecutive(nums []int) int {
	seen := make(map[int]bool, len(nums))
	best := 0
	
	for _, num := range nums {
		seen[num] = true
	}

	for x := range seen {
		if seen[x-1] {
			continue
		}

		length := 1
		cur := x

		for seen[cur+1] {
			cur++
			length++
		}
		
		best = max(best, length)
	}

	return best
}
