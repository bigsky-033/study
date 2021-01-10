package p_00401_00500

// 413. Arithmetic Slices, https://leetcode.com/problems/arithmetic-slices/

func numberOfArithmeticSlices(A []int) int {
	n := 0
	end, start := 2, 0
	for ; end < len(A); end++ {
		if A[end]-A[end-1] != A[end-1]-A[end-2] {
			n += helper(end-1, start)
			start = end - 1
		}
	}
	n += helper(end-1, start)
	return n
}

func helper(end int, start int) int {
	n := end - start + 1 - 2
	return n * (n + 1) / 2
}
