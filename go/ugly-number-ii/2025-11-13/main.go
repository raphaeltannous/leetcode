import "container/heap"

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func nthUglyNumber(n int) int {
	h := &minHeap{1}
	hashMap := make(map[int]bool)
	heap.Init(h)

	allowed := []int{2, 3, 5}
	currentUgly := 1
	for i := 0; i < n; i++ {
		currentUgly = heap.Pop(h).(int)

		for _, j := range allowed {
			if value := currentUgly * j; !hashMap[value] {
				heap.Push(h, value)
				hashMap[value] = true
			}
		}
	}

	return currentUgly
}
