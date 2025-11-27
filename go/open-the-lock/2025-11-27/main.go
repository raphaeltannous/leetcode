
import "math"

func openLock(deadends []string, target string) int {
	deadendsMap := make(map[string]bool)
	for _, deadend := range deadends {
		deadendsMap[deadend] = true
	}

	result := math.MaxInt
	queue := []string{"0000"}
	rotation := 0
	for len(queue) > 0 {
		lockCount := len(queue)

		for range lockCount {
			lock := queue[0]
			queue = queue[1:]

			if deadendsMap[lock] {
				continue
			}

			if lock == target {
				result = min(result, rotation)
				break
			}

			deadendsMap[lock] = true

			for i := range 4 {
				queue = append(queue, addOne(lock, i), subOne(lock, i))
			}
		}

		rotation++
	}

	if result == math.MaxInt {
		return -1
	}

	return result
}

func addOne(lock string, position int) string {
	number := lock[position]

	if number == '9' {
		number = '0'
	} else {
		number += 1
	}

	return lock[:position] + string(number) + lock[position+1:]
}

func subOne(lock string, position int) string {
	number := lock[position]

	if number == '0' {
		number = '9'
	} else {
		number -= 1
	}

	return lock[:position] + string(number) + lock[position+1:]
}
