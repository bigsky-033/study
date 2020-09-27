package p_00601_00700

// 622. Design Circular Queue, https://leetcode.com/problems/design-circular-queue/

type MyCircularQueue struct {
	array    []int
	capacity int
	size     int
	head     int
	tail     int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		array:    make([]int, k),
		capacity: k,
		size:     0,
		head:     0,
		tail:     0,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.array[q.tail] = value
	q.tail = (q.tail + 1) % q.capacity
	q.size++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.head = (q.head + 1) % q.capacity
	q.size--
	return true
}

/** Get the front item from the queue. */
func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.array[q.head]
}

/** Get the last item from the queue. */
func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.array[(q.tail+q.capacity-1)%q.capacity]
}

/** Checks whether the circular queue is empty or not. */
func (q *MyCircularQueue) IsEmpty() bool {
	return q.size == 0
}

/** Checks whether the circular queue is full or not. */
func (q *MyCircularQueue) IsFull() bool {
	return q.capacity == q.size
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
