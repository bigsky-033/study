package main

import "strconv"

func openLock(deadends []string, target string) int {
	empty := struct{}{}

	counter := 0

	deadendsSet := make(map[string]struct{})
	visited := make(map[string]struct{})
	for _, v := range deadends {
		deadendsSet[v] = empty
	}

	var queue []string
	queue = append(queue, "0000")
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			wheel := queue[0]
			queue = queue[1:]
			if wheel == target {
				return counter
			}
			if _, ok := visited[wheel]; ok {
				continue
			}
			if _, ok := deadendsSet[wheel]; ok {
				continue
			}
			visited[wheel] = empty

			for j := 0; j < len(wheel); j++ {
				k := int(wheel[j] - '0')
				up := (k + 1) % 10
				down := (k + 9) % 10
				queue = append(queue, wheel[0:j]+strconv.Itoa(up)+wheel[j+1:])
				queue = append(queue, wheel[0:j]+strconv.Itoa(down)+wheel[j+1:])
			}
		}
		counter++
	}
	return -1
}
