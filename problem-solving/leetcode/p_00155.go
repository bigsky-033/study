package main

// 155. Min Stack, https://leetcode.com/problems/min-stack/

type MinStack struct {
	storage []int
	min     int // min points to latest minimum value. min value is not stored on storage
}

const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
const MaxInt = 1<<(UintSize-1) - 1

/** initialize your data structure here. */
func Constructor() MinStack {
	var storage []int
	return MinStack{storage: storage, min: MaxInt}
}

func (s *MinStack) Push(x int) {
	if x <= s.min {
		s.storage = append(s.storage, s.min)
		s.min = x
	}
	s.storage = append(s.storage, x)
}

func (s *MinStack) Pop() {
	if len(s.storage) < 1 {
		return
	}
	targetIdx := len(s.storage) - 1
	popped := s.storage[targetIdx]
	s.storage = s.storage[:targetIdx]

	if popped == s.min {
		targetIdx = len(s.storage) - 1
		popped = s.storage[targetIdx]
		s.storage = s.storage[:targetIdx]
		s.min = popped
	}
}

func (s *MinStack) Top() int {
	return s.storage[len(s.storage)-1]
}

func (s *MinStack) GetMin() int {
	return s.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
