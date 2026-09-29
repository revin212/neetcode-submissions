type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func isValid(s string) bool {
    pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := &Stack[rune]{}

	for _, ch := range s {
		if opener, isCloser := pairs[ch]; isCloser {
			top, ok := stack.Pop()
			if !ok || top != opener {
				return false
			}
		} else {
			stack.Push(ch)
		}
	}

	return stack.IsEmpty()
}
