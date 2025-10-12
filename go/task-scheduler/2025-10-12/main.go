import "container/list"

type maxHeap struct {
	heap []int
}

func newMaxHeap(tasks []byte) maxHeap {
	counter := make(map[int]int)
	for _, task := range tasks {
		counter[int(task)]++
	}

	h := maxHeap{
		heap: make([]int, 0, len(counter)),
	}

	for _, v := range counter {
		h.Push(v)
	}

	return h
}

func (mh *maxHeap) Push(taskCount int) {
	mh.heap = append(mh.heap, taskCount)
	mh.heapifyUp()
}

func (mh *maxHeap) heapifyUp() {
	index := mh.Len() - 1
	parent := mh.parent(index)

	for mh.heap[parent] < mh.heap[index] {
		mh.swap(index, parent)
		index = parent
		parent = mh.parent(index)
	}
}

func (mh *maxHeap) Pop() int {
	top := mh.heap[0]

	mh.heap[0] = mh.heap[mh.Len()-1]
	mh.heap = mh.heap[:mh.Len()-1]

	mh.heapifyDown()

	return top
}

func (mh *maxHeap) heapifyDown() {
	index, lastIndex := 0, mh.Len()-1
	left, right := mh.left(index), mh.right(index)

	for left <= lastIndex {
		childToCompare := 0
		if left == lastIndex { // no right
			childToCompare = left
		} else if mh.heap[left] > mh.heap[right] { // left > right
			childToCompare = left
		} else { // right > left
			childToCompare = right
		}

		if mh.heap[childToCompare] > mh.heap[index] {
			mh.swap(childToCompare, index)
			index = childToCompare
			left, right = mh.left(index), mh.right(index)
		} else {
			return
		}
	}
}

func (mh *maxHeap) swap(i, j int) {
	mh.heap[i], mh.heap[j] = mh.heap[j], mh.heap[i]
}

func (mh *maxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (mh *maxHeap) left(i int) int {
	return (i * 2) + 1
}

func (mh *maxHeap) right(i int) int {
	return (i * 2) + 2
}

func (mh *maxHeap) Len() int {
	return len(mh.heap)
}

func leastInterval(tasks []byte, n int) int {
	interval := 0
	queue := list.New()
	h := newMaxHeap(tasks)

	for queue.Len() != 0 || h.Len() != 0 {
		if h.Len() != 0 {
			top := h.Pop()

			if top-1 > 0 {
				queue.PushBack([2]int{top - 1, interval + n})
			}
		}

		if queue.Len() != 0 {
			front := queue.Front()

			if front.Value.([2]int)[1] == interval {
				h.Push(front.Value.([2]int)[0])
				queue.Remove(front)
			}
		}

		interval++
	}

	return interval
}
