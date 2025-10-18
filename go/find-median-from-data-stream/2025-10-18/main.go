type MedianFinder struct {
	miHeap *minHeap
	maHeap *maxHeap
}

// maxHeap
type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func (m *maxHeap) Top() any {
	return (*m)[0]
}

// minHeap
type minHeap []int

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func (m *minHeap) Top() any {
	return (*m)[0]
}

func Constructor() MedianFinder {
	miHeap := &minHeap{}
	heap.Init(miHeap)

	maHeap := &maxHeap{}
	heap.Init(maHeap)

	return MedianFinder{
		miHeap: miHeap,
		maHeap: maHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.miHeap, num)

	if this.maHeap.Len() > 0 && this.miHeap.Top().(int) < this.maHeap.Top().(int) {
		miTop := heap.Pop(this.miHeap).(int)
		maTop := heap.Pop(this.maHeap).(int)

		heap.Push(this.miHeap, maTop)
		heap.Push(this.maHeap, miTop)
	}

	if this.miHeap.Len()-this.maHeap.Len() > 1 {
		heap.Push(this.maHeap, heap.Pop(this.miHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	var result float64

	length := this.miHeap.Len() + this.maHeap.Len()
	if length%2 != 0 {
		result = float64(this.miHeap.Top().(int))
	} else {
		result = float64(this.miHeap.Top().(int)+this.maHeap.Top().(int)) / 2.0
	}

	return result
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
