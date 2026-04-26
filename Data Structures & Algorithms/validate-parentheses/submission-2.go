type Stack struct{
	data []rune
}

func (s *Stack) Push(bracket rune) {
	s.data = append(s.data, bracket)
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	lastIndex := len(s.data) - 1
	value := s.data[lastIndex]
	s.data = s.data[:lastIndex]

	return value, true
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func isValid(s string) bool {
	stack := Stack{}

    brackets := map[rune]rune {
		'}': '{',
		']': '[',
		')': '(',
	}
	
	for _, elem := range s {
		if elem == '{' || elem == '[' || elem == '(' {
			stack.Push(elem)
			continue
		}

		if elem == '}' || elem == ']' || elem == ')' {
			popped, ok := stack.Pop()
			if !ok || popped != brackets[elem] {
				return false
			}
		}
	}

	return stack.Empty()
}
