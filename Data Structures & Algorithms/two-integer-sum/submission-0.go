func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	var res []int

    for i, num := range nums {
		need := target - num
		j, ok := m[need]
		if ok {
			if i < j {
				res = []int{i, j}
			} else {
				res = []int{j, i}
			}
			return res
		} else {
			m[num] = i
		}
	}
	return res
}
