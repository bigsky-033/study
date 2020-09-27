package p_00301_00400

// 374. Guess Number Higher or Lower, https://leetcode.com/problems/guess-number-higher-or-lower/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	lo, hi := 1, n

	for lo < hi {
		// with upper/right mid
		// mid := lo + (hi - lo + 1)/2
		// res := guess(mid)
		// if res == -1 {
		//     hi = mid - 1
		// } else {
		//     lo = mid
		// }

		// with left/lower mid
		mid := lo + (hi-lo)/2
		res := guess(mid)
		if res == 1 {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
