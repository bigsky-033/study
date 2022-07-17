package p01342numberofstepstoreduceanumbertozero

// numberOfSteps iterative version
func numberOfSteps(num int) int {
	steps := 0
	for num != 0 {
		if isEven(num) {
			num = num / 2
		} else {
			num -= 1
		}
		steps++
	}
	return steps
}

func isEven(num int) bool {
	return num%2 == 0
}

// numberOfSteps iterative version with bitwise operations
// func numberOfSteps(num int) int {
// 	steps := 0
// 	for num > 0 {
// 		if (num & 1) == 0 {
// 			num >>= 1
// 		} else {
// 			num--
// 		}
// 		steps++
// 	}
// 	return steps
// }

// numberOfSteps recursive version
// func numberOfSteps(num int) int {
// 	if num == 1 {
// 		return 1
// 	}
// 	if num == 2 {
// 		return 2
// 	}
// 	if isEven(num) {
// 		return numberOfSteps(num/2) + 1
// 	} else {
// 		return numberOfSteps(num-1) + 1
// 	}
// }
