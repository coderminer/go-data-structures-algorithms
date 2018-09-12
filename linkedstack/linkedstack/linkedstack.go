package linkedstack

import (
	"fmt"
)

type Node struct {
	val  interface{}
	next *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) Push(val interface{}) {
	node := &Node{val, nil}
	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}
	s.size += 1
}

func (s *Stack) Pop() interface{} {
	if s.head == nil {
		return nil
	} else {
		h := s.head
		s.head = s.head.next
		s.size -= 1
		return h.val
	}

}

func (s *Stack) Top() interface{} {
	if s.head == nil {
		return nil
	} else {
		return s.head.val
	}
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) String() {
	cur := s.head
	for cur != nil {
		fmt.Print(cur.val, "->")
		cur = cur.next
	}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) IsBalancedParenthesis(exp string) bool {
	for _, ch := range exp {
		switch ch {
		case '{', '[', '(':
			s.Push(ch)
		case '}':
			val := s.Pop()
			if val != '{' {
				return false
			}
		case ']':
			val := s.Pop()
			if val != '[' {
				return false
			}
		case ')':
			val := s.Pop()
			if val != '(' {
				return false
			}
		}

	}

	return s.IsEmpty()
}
