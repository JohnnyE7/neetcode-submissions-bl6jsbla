func isAnagram(s string, t string) bool {
	var cnt [26]int
	if len(s) != len(t) { return false }

	for i := 0; i < len(s); i++ {
		cnt[s[i] - 'a']++
		cnt[t[i] - 'a']--
	}

	for _, ch := range cnt {
		if ch != 0 { return false }
	}

	return true
}
