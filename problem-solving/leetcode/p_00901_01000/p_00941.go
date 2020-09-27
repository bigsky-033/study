package p_00901_01000

// 941. Valid Mountain Array, https://leetcode.com/problems/valid-mountain-array/

func validMountainArray(A []int) bool {
	if len(A) < 3 || A[0] > A[1] {
		return false
	}
	i := 0
	for ; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			return false
		} else if A[i] > A[i+1] {
			break
		}
	}
	if i == len(A)-1 {
		return false
	}
	i++
	for ; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			return false
		} else if A[i] < A[i+1] {
			return false
		}
	}
	return true
}
