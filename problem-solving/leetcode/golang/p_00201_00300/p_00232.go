package p_00201_00300

type MyQueue struct {
	stack1 []int
	stack2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	for len(q.stack1) > 0 {
		v := q.stack1[len(q.stack1)-1]
		q.stack1 = q.stack1[:len(q.stack1)-1]
		q.stack2 = append(q.stack2, v)
	}
	q.stack1 = append(q.stack1, x)
	for len(q.stack2) > 0 {
		v := q.stack2[len(q.stack2)-1]
		q.stack2 = q.stack2[:len(q.stack2)-1]
		q.stack1 = append(q.stack1, v)
	}
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	v := q.stack1[len(q.stack1)-1]
	q.stack1 = q.stack1[:len(q.stack1)-1]
	return v
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	return q.stack1[len(q.stack1)-1]
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return len(q.stack1) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
