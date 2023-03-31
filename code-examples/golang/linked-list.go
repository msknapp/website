package golang

type node[T any] struct {
	Value    T
	Next     *node[T]
	Previous *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		Value: value,
	}
}

type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Push(value T) {
	n := newNode(value)
	l.size++
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	n.Previous = l.tail
	l.tail.Next = n
	l.tail = n
}

func (l *LinkedList[T]) PushToFront(value T) {
	n := newNode(value)
	l.size++
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	l.head.Previous = n
	n.Next = l.head
	l.head = n
}

func (l *LinkedList[T]) Pop() T {
	var out T
	if l.tail == nil {
		return out
	}
	out = l.tail.Value
	l.tail = l.tail.Previous
	if l.tail != nil {
		l.tail.Next = nil
	}
	l.size--
	return out
}

func (l *LinkedList[T]) PopFront() T {
	var out T
	if l.head == nil {
		return out
	}
	l.size--
	out = l.head.Value
	l.head = l.head.Next
	if l.head != nil {
		l.head.Previous = nil
	}
	return out
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) Peek() T {
	if l.tail == nil {
		var out T
		return out
	}
	return l.tail.Value
}

func (l *LinkedList[T]) PeekFront() T {
	if l.head == nil {
		var out T
		return out
	}
	return l.head.Value
}
