package stack

type Stack[T any] struct {
	Elements []T
}

func Build[T any](size, cap int) *Stack[T] {
	elems := make([]T, size, cap)
	return &Stack[T]{
		Elements: elems,
	}
}

func (s *Stack[T]) Push(elem T) {
	s.Elements = append(s.Elements, elem)
}

func (s *Stack[T]) Pop() (T, bool) {
	var elem T
	if len(s.Elements) > 0 {
		elem = s.Elements[len(s.Elements)-1]
		s.Elements = s.Elements[:len(s.Elements)-1]
		return elem, true
	}
	return elem, false
}
