package p_00801_00900

// 832. Flipping an Image, https://leetcode.com/problems/flipping-an-image/

func flipAndInvertImage(A [][]int) [][]int {
	for _, r := range A {
		left, right := 0, len(r)-1
		for left < right {
			tmp := r[left] ^ 1
			r[left] = r[right] ^ 1
			r[right] = tmp

			left++
			right--
		}
		if len(r)%2 != 0 {
			r[len(r)/2] = r[len(r)/2] ^ 1
		}
	}
	return A
}
