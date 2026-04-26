type Stack struct{
	Data []rune
}

func (s *Stack) Push(bracket rune) {
	s.Data = append(s.Data, bracket)
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.Data) == 0 {
		return 0, false
	}

	lastIndex := len(s.Data) - 1
	value := s.Data[lastIndex]
	s.Data = s.Data[:lastIndex]

	return value, true
}

func (s *Stack) Empty() bool {
	return len(s.Data) == 0
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
