import "container/heap"

type minHeap [][]int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	if h[i][1] == h[j][1] {
		return h[i][2] < h[j][2]
	}

	return h[i][1] < h[j][1]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h minHeap) Top() any {
	return h[h.Len()-1]
}

func carPooling(trips [][]int, capacity int) bool {
	h := minHeap(trips)
	heap.Init(&h)

	capacityLeft := capacity
	for h.Len() > 0 {
		top := h[0]

		fmt.Println(top, capacityLeft)
		if top[0] <= capacityLeft {
			capacityLeft -= top[0]
			heap.Pop(&h)
			if top[2] != -1 {
				heap.Push(&h, []int{-top[0], top[2], -1})
			}
		} else {
			return false
		}
	}

	return true
}
