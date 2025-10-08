type maxHeap struct {
	heap []int
}

func NewMaxHeap(nums []int) maxHeap {
	heap := maxHeap{
		heap: make([]int, 0, len(nums)),
	}

	for _, num := range nums {
		heap.Insert(num)
	}

	return heap
}

func (mh *maxHeap) Insert(value int) {
	mh.heap = append(mh.heap, value)
	mh.heapifyUp()
}

func (mh *maxHeap) GetTop() int {
	if len(mh.heap) == 0 {
		return 0
	}

	value := mh.heap[0]
	mh.heap[0] = mh.heap[len(mh.heap)-1]
	mh.heap = mh.heap[:len(mh.heap)-1]

	mh.heapifyDown()
	return value
}

func (mh *maxHeap) heapifyUp() {
	index := len(mh.heap) - 1

	for mh.heap[mh.parent(index)] < mh.heap[index] {
		mh.swap(mh.parent(index), index)
		index = mh.parent(index)
	}
}

func (mh *maxHeap) heapifyDown() {
	index := 0
	left, right := mh.left(index), mh.right(index)

	for left <= len(mh.heap)-1 {
		childToCompare := 0
		if left == len(mh.heap)-1 {
			childToCompare = left
		} else if mh.heap[left] > mh.heap[right] {
			childToCompare = left
		} else {
			childToCompare = right
		}

		if mh.heap[index] < mh.heap[childToCompare] {
			mh.swap(index, childToCompare)
			index = childToCompare
			left, right = mh.left(index), mh.right(index)
		} else {
			break
		}
	}
}

func (mh *maxHeap) swap(index1, index2 int) {
	mh.heap[index1], mh.heap[index2] = mh.heap[index2], mh.heap[index1]
}

func (mh *maxHeap) parent(index int) int {
	return (index - 1) / 2
}

func (mh *maxHeap) left(index int) int {
	return (index * 2) + 1
}

func (mh *maxHeap) right(index int) int {
	return (index * 2) + 2
}

func (mh *maxHeap) Len() int {
	return len(mh.heap)
}

func lastStoneWeight(stones []int) int {
	heap := NewMaxHeap(stones)

	for heap.Len() > 1 {
		y := heap.GetTop()
		x := heap.GetTop()

		if (y - x) > 0 {
			heap.Insert(y - x)
		}
	}

	return heap.GetTop()
}
