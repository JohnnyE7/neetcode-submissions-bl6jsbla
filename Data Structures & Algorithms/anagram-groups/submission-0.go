func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	var res [][]string
	
	for _, str := range strs {
		key := buildKey(str)
		m[key] = append(m[key], str)
	}

	for _, str := range m {
		res = append(res, str)
	}

	return res
}

func buildKey(str string) [26]int {
	var key [26]int
	for i := 0; i < len(str); i++ {
		key[str[i]-'a']++
	}
	return key
}
