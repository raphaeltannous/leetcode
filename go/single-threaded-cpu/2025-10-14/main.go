type minHeap struct {
	heap [][3]int
}

func newMinHeap(tasks [][]int) minHeap {
	heap := minHeap{
		heap: make([][3]int, 0, len(tasks)),
	}

	for i, task := range tasks {
		heap.Push([3]int{i, task[0], task[1]})
	}

	return heap
}

func (mh *minHeap) Push(task [3]int) {
	mh.heap = append(mh.heap, task)
	mh.heapifyUp(mh.Less)
}

func (mh *minHeap) PushShortest(task [3]int) {
	mh.heap = append(mh.heap, task)
	mh.heapifyUp(mh.LessShortest)
}

func (mh *minHeap) heapifyUp(compareFunc func(i, j int) bool) {
	index := mh.Len() - 1
	parent := mh.parent(index)

	for compareFunc(index, parent) {
		mh.swap(parent, index)
		index = parent
		parent = mh.parent(index)
	}
}

func (mh *minHeap) PopShortest() [3]int {
	top := mh.heap[0]

	mh.heap[0] = mh.heap[mh.Len()-1]
	mh.heap = mh.heap[:mh.Len()-1]

	mh.heapifyDown(mh.LessShortest)

	return top
}

func (mh *minHeap) Pop() [3]int {
	top := mh.heap[0]

	mh.heap[0] = mh.heap[mh.Len()-1]
	mh.heap = mh.heap[:mh.Len()-1]

	mh.heapifyDown(mh.Less)

	return top
}

func (mh *minHeap) heapifyDown(compareFunc func(i, j int) bool) {
	index, lastIndex := 0, mh.Len()-1
	left, right := mh.left(index), mh.right(index)

	for left <= lastIndex {
		childToCompare := 0
		if left == lastIndex {
			childToCompare = left
		} else if compareFunc(left, right) {
			childToCompare = left
		} else {
			childToCompare = right
		}

		if compareFunc(childToCompare, index) {
			mh.swap(childToCompare, index)
			index = childToCompare
			left, right = mh.left(index), mh.right(index)
		} else {
			return
		}
	}
}

func (mh *minHeap) swap(i, j int) {
	mh.heap[i], mh.heap[j] = mh.heap[j], mh.heap[i]
}

func (mh *minHeap) Len() int {
	return len(mh.heap)
}

func (mh *minHeap) parent(i int) int {
	return (i - 1) / 2
}

func (mh *minHeap) left(i int) int {
	return (i * 2) + 1
}

func (mh *minHeap) right(i int) int {
	return (i * 2) + 2
}

func (mh *minHeap) LessShortest(i, j int) bool {
	if mh.heap[i][2] == mh.heap[j][2] {
		return mh.heap[i][0] < mh.heap[j][0]
	}

	return mh.heap[i][2] < mh.heap[j][2]
}

func (mh *minHeap) Less(i, j int) bool {
	// enqueue time
	if mh.heap[i][1] < mh.heap[j][1] {
		return true
	} else if mh.heap[i][1] == mh.heap[j][1] {
		// process time
		if mh.heap[i][2] < mh.heap[j][2] {
			return true
		} else if mh.heap[i][2] == mh.heap[j][2] {
			// index
			if mh.heap[i][0] < mh.heap[j][0] {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

func (mh *minHeap) Top() [3]int {
	return mh.heap[0]
}

func getOrder(tasks [][]int) []int {
	result := make([]int, 0, len(tasks))
	heap := newMinHeap(tasks)
	time := heap.Top()[1]
	aHeap := newMinHeap([][]int{})

	for heap.Len() != 0 || aHeap.Len() != 0 {
		if heap.Len() != 0 {
			top := heap.Top()

			for top[1] <= time {
				aHeap.PushShortest(heap.Pop())
				if heap.Len() != 0 {
					top = heap.Top()
				} else {
					break
				}
			}
		}

		aTop := aHeap.PopShortest()
		time += aTop[2]
		result = append(result, aTop[0])

		if aHeap.Len() == 0 && heap.Len() != 0 {
			fmt.Println("prev", time)
			newTime := heap.Top()[1]
			if newTime > time {
				time = newTime
			}
			fmt.Println("curr", time)
		}
	}

	return result
}
