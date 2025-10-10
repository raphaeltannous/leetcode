type minKHeap struct {
	k    int
	heap []int
}

func (mh *minKHeap) Push(value int) {
	if mh.Len() == mh.k && mh.heap[0] > value {
		return
	}

	if mh.Len() == mh.k {
		mh.Pop()
	}

	mh.heap = append(mh.heap, value)
	mh.heapifyUp()
}

func (mh *minKHeap) Pop() int {
	value := mh.heap[0]
	mh.heap[0] = mh.heap[mh.Len()-1]
	mh.heap = mh.heap[:mh.Len()-1]
	mh.heapifyDown()
	return value
}

func (mh *minKHeap) heapifyDown() {
	index, lastIndex := 0, mh.Len()-1
	left, right := mh.left(index), mh.right(index)

	for left <= lastIndex {
		childToCompare := 0
		if left == lastIndex {
			childToCompare = left
		} else if mh.heap[left] < mh.heap[right] {
			childToCompare = left
		} else {
			childToCompare = right
		}

		if mh.heap[index] > mh.heap[childToCompare] {
			mh.swap(index, childToCompare)
			index = childToCompare
			left, right = mh.left(index), mh.right(index)
		} else {
			return
		}
	}
}

func (mh *minKHeap) heapifyUp() {
	index := mh.Len() - 1
	parentIndex := mh.parent(index)
	for mh.heap[parentIndex] >= mh.heap[index] {
		if index == 0 {
			break
		}

		mh.swap(parentIndex, index)
		index = parentIndex
		parentIndex = mh.parent(index)
	}
}

func (mh *minKHeap) Len() int {
	return len(mh.heap)
}

func (mh *minKHeap) swap(i, j int) {
	mh.heap[i], mh.heap[j] = mh.heap[j], mh.heap[i]
}

func (mh *minKHeap) parent(i int) int {
	return (i - 1) / 2
}

func (mh *minKHeap) left(i int) int {
	return (i * 2) + 1
}

func (mh *minKHeap) right(i int) int {
	return (i * 2) + 2
}

func findKthLargest(nums []int, k int) int {
	mh := minKHeap{
		k:    k,
		heap: make([]int, 0, k),
	}

	for _, num := range nums {
		mh.Push(num)
	}

	return mh.Pop()
}
