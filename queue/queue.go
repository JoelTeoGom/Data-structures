package queue

type Queue[T any] struct {
	Elements []T
}

func Build[T any](size, cap int) *Queue[T] {
	elem := make([]T, size, cap)
	return &Queue[T]{
		Elements: elem,
	}
}

func (q *Queue[T]) Queue(element T) {
	q.Elements = append(q.Elements, element)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var elem T
	if len(q.Elements) > 0 {
		elem = q.Elements[0]
		q.Elements = q.Elements[1:len(q.Elements)]
		return elem, true
	}
	return elem, false
}
