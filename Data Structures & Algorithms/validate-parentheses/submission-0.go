type Stack struct{
	Data []byte
}

func (s *Stack) Push(bracket byte) {
	s.Data = append(s.Data, bracket)
}

func (s *Stack) Pop() (byte, bool) {
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

    brackets := map[byte]byte {
		'}': '{',
		']': '[',
		')': '(',
	}
	
	for _, elem := range s {
		ebyte := byte(elem)
		if ebyte == '{' || ebyte == '[' || ebyte == '(' {
			stack.Push(ebyte)
			continue
		}

		if ebyte == '}' || ebyte == ']' || ebyte == ')' {
			popped, ok := stack.Pop()
			if !ok || popped != brackets[ebyte] {
				return false
			}
		}
	}

	return stack.Empty()
}
