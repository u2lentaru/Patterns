package main

import "errors"

// Stack structure
type Stack struct {
	head *node
	tail *node
	len  int
}

type node struct {
	data interface{}
	next *node
	prev *node
}

// Push method
func (l *Stack) Push(data interface{}) {
	n := &node{
		data: data,
	}

	if l.head == nil && l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}

	l.len++
}

/*func (l *Stack) get(i int) (*node, error) {
	if i < 0 || i >= l.len {
		return nil, errors.New("out of range")
	}

	index := 0
	current := l.head

	for index != i {
		current = current.next
		index++
	}

	return current, nil
}*/

// Len method
func (l *Stack) Len() int {
	return l.len
}

// Get method
/*func (l *Stack) Get(i int) (interface{}, error) {
	node, err := l.get(i)
	if err != nil {
		return nil, err
	}
	return node.data, nil
}*/

// Pop method
//func (l *Stack) Pop(i int) error {
func (l *Stack) Pop() (interface{}, error) {
	if l.head == nil && l.tail == nil {
		return nil, errors.New("stack is empty")
	}

	if l.head == l.tail {
		d := l.head.data
		l.head = nil
		l.tail = nil
		l.len--
		return d, nil
	}

	d := l.head.data
	l.head = l.head.next
	l.len--
	return d, nil

}

// ToArray method
func (l *Stack) ToArray() []interface{} {
	arr := make([]interface{}, 0, l.len)

	current := l.head
	for current != nil {
		arr = append(arr, current.data)
		current = current.next
	}

	return arr
}
