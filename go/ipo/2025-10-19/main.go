import "container/heap"

type heapInterface interface {
	heap.Interface
	Top() any
}

type minHeap [][2]int

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	// [2]int{capital, profit}
	if h[i][0] < h[j][0] {
		return true
	} else if h[i][0] > h[j][0] {
		return false
	}

	return h[i][1] > h[j][1]
}
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h minHeap) Top() any {
	return h[0]
}

type maxHeap []int

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	// we are not interested in the capital
	return h[i] > h[j]
}
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h maxHeap) Top() any {
	return h[0]
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	miHeap := &minHeap{}
	heap.Init(miHeap)

	for x, p := range profits {
		heap.Push(miHeap, [2]int{capital[x], p})
	}

	currentCapital := w

	maHeap := &maxHeap{}
	heap.Init(maHeap)

	for range k {
		// populate with respect to capital
		populateUntilCapitale(miHeap, maHeap, currentCapital)

		if maHeap.Len() > 0 {
			currentCapital += heap.Pop(maHeap).(int)
		} else {
			break
		}
	}

	return currentCapital
}

func populateUntilCapitale(miHeap, maHeap heapInterface, w int) {
	if miHeap.Len() == 0 {
		return
	}

	top := miHeap.Top().([2]int)
	for top[0] <= w {
		heap.Push(maHeap, heap.Pop(miHeap).([2]int)[1])
		if miHeap.Len() > 0 {
			top = miHeap.Top().([2]int)
		} else {
			break
		}
	}
}
