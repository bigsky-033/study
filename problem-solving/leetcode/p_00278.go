package main

// 278. First Bad Version, https://leetcode.com/problems/first-bad-version/

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if !isBadVersion(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
