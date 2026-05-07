type Stack struct{
	buf []int
}

func (s *Stack) operation(char string) int {
	var res int

	left := s.buf[len(s.buf)-2]
	right := s.buf[len(s.buf)-1]

	switch char {
	case "+": 
		res = left + right
	case "-": 
		res = left - right
	case "*":
		res = left * right
	case "/":
		res = left / right
	}

	return res
}

func (s *Stack) clear() {
	s.buf = s.buf[:len(s.buf) - 2]
}

func evalRPN(tokens []string) int {
	s := Stack{}
	var res int

	for _, char := range tokens {
		operand, err := strconv.Atoi(char)

		if err == nil {
			s.buf = append(s.buf, operand)
		} else {
			res = s.operation(char)
			s.clear()
			s.buf = append(s.buf, res)
		}
	}

	return s.buf[0]
}