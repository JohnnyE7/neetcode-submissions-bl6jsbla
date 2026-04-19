type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var res strings.Builder

	for _, str := range strs {
		res.WriteString(strconv.Itoa(len(str)))
		res.WriteByte('#')
		res.WriteString(str)
	}

	return res.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)

	for i := 0; i < len(encoded); {
		j := i
		for encoded[j] != '#' {
			j++
		}
		n, _ := strconv.Atoi(encoded[i:j])
		j++
		res = append(res, encoded[j:j+n])
		i = j + n
	}

	return res
}
