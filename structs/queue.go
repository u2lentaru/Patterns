package main

import (
	"errors"
)

// Queue structure
type Queue struct {
	head *node
	tail *node
	len  int
}

/*type node struct {
	data interface{}
	next *node
	prev *node
}*/

// Push method
func (l *Queue) Push(data interface{}) {
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

// Len method
func (l *Queue) Len() int {
	return l.len
}

// Pop method
func (l *Queue) Pop() (interface{}, error) {
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
	l.head.prev = nil
	l.len--
	return d, nil
}

// ToArray method
func (l *Queue) ToArray() []interface{} {
	arr := make([]interface{}, 0, l.len)

	current := l.head
	for current != nil {
		arr = append(arr, current.data)
		current = current.next
	}
	return arr
}
