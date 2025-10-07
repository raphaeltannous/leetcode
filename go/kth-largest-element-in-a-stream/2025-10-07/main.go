type KthLargest struct {
	k       int
	minHeap []int
}

func Constructor(k int, nums []int) KthLargest {
	s := KthLargest{
		k:       k - 1,
		minHeap: make([]int, 0, k),
	}

	for _, num := range nums {
		s.Add(num)
	}

	return s
}

func (this *KthLargest) Add(val int) int {
	if len(this.minHeap) < this.k+1 {
		this.minHeap = append(this.minHeap, val)
		this.heapifyUp()
	} else if this.minHeap[0] < val {
		this.removeTop()
		this.Add(val)
	}

	return this.minHeap[0]
}

func (this *KthLargest) heapifyUp() {
	index := len(this.minHeap) - 1
	for this.minHeap[this.parent(index)] > this.minHeap[index] {
		this.minHeap[this.parent(index)], this.minHeap[index] = this.minHeap[index], this.minHeap[this.parent(index)]
		index = this.parent(index)
	}
}

func (this *KthLargest) removeTop() {
	this.minHeap[0] = this.minHeap[len(this.minHeap)-1]
	this.minHeap = this.minHeap[:len(this.minHeap)-1]

	this.heapifyDown()
}

func (this *KthLargest) heapifyDown() {
	index := 0
	lastIndex := len(this.minHeap) - 1
	left, right := this.left(index), this.right(index)
	var childToCompare int
	for left <= lastIndex {
		if left == lastIndex {
			childToCompare = left
		} else if this.minHeap[left] < this.minHeap[right] {
			childToCompare = left
		} else {
			childToCompare = right
		}

		if this.minHeap[index] > this.minHeap[childToCompare] {
			this.minHeap[index], this.minHeap[childToCompare] = this.minHeap[childToCompare], this.minHeap[index]
			index = childToCompare
			left, right = this.left(index), this.right(index)
		} else {
			return
		}
	}
}

func (this *KthLargest) parent(index int) int {
	return (index - 1) / 2
}

func (this *KthLargest) left(index int) int {
	return (index * 2) + 1
}

func (this *KthLargest) right(index int) int {
	return (index * 2) + 2
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
