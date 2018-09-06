package main

import (
	"fmt"
)

/*
	使用slice定义栈
*/
type Stack struct {
	elements []int
}

func (s *Stack) Push(ele int) {
	s.elements = append(s.elements, ele)
}

func (s *Stack) Top() int {
	if len(s.elements) > 0 {
		return s.elements[0]
	}
	return -1
}

func (s *Stack) Pop() int {
	if len(s.elements) > 0 {
		top := s.elements[0]
		s.elements = s.elements[1:]

		return top
	}
	return -1
}

func (s *Stack) Size() int {
	return len(s.elements)
}

type StackMin struct {
	mData *Stack
	mMin  *Stack
}

func (s *StackMin) MinPush(val int) {
	s.mData.Push(val)
	if s.mMin.Size() == 0 || val < s.mMin.Top() {
		s.mMin.Push(val)
	} else {
		s.mMin.Push(s.mMin.Top())
	}
}

func (s *StackMin) MinPop() {
	s.mData.Pop()
	s.mMin.Pop()
}

func (s *StackMin) Min() int {
	return s.mMin.Top()
}

/*
	定义栈，实现一个能够得到的最小的元素，在该栈中，min push pop 的时间复杂度都是O(1)
	需要使用一个辅助栈
*/

func main() {
	s := StackMin{}
	s.mData = &Stack{}
	s.mMin = &Stack{}
	
	s.MinPush(3)
	s.MinPush(4)
	s.MinPush(2)
	s.MinPush(5)

	fmt.Println(s.Min())
}
