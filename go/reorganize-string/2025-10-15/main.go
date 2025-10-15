import "container/list"
import "strings"

type maxHeap struct {
	heap []*element
}

func newMaxHeap(s string) maxHeap {
	counter := make(map[rune]int)

	for _, letter := range s {
		counter[letter]++
	}

	heap := maxHeap{
		heap: make([]*element, 0, len(counter)),
	}

	for letter, frequency := range counter {
		heap.Push(&element{letter: letter, count: frequency})
	}

	return heap
}

func (mh *maxHeap) Push(e *element) {
	mh.heap = append(mh.heap, e)
	mh.heapifyUp()
}

func (mh *maxHeap) heapifyUp() {
	index := mh.Len() - 1
	parent := mh.parent(index)

	for mh.Less(parent, index) {
		mh.swap(index, parent)
		index = parent
		parent = mh.parent(index)
	}
}

func (mh *maxHeap) Pop() *element {
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
		var childToCompare int
		if left == lastIndex {
			childToCompare = left
		} else if mh.Less(right, left) {
			childToCompare = left
		} else {
			childToCompare = right
		}

		if mh.Less(index, childToCompare) {
			mh.swap(index, childToCompare)
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

func (mh *maxHeap) left(i int) int {
	return (i * 2) + 1
}

func (mh *maxHeap) right(i int) int {
	return (i * 2) + 2
}

func (mh *maxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (mh *maxHeap) Len() int {
	return len(mh.heap)
}

func (mh *maxHeap) Top() *element {
	if mh.Len() != 0 {
		return mh.heap[0]
	}

	return nil
}

func (mh *maxHeap) Less(i, j int) bool {
	return mh.heap[i].count < mh.heap[j].count
}

type element struct {
	letter rune
	count  int
}

func reorganizeString(s string) string {
	heap := newMaxHeap(s)
	queue := list.New()

	var result strings.Builder
	var lastLetter rune
	x := 1

	for heap.Len() != 0 || queue.Len() != 0 {
		if heap.Len() != 0 && x%2 != 0 {
			if top := heap.Top(); top.letter != lastLetter {
				result.WriteRune(top.letter)
				fmt.Println(string(top.letter))
				lastLetter = top.letter

				removedTop := heap.Pop()
				removedTop.count--
				if removedTop.count > 0 {
					queue.PushBack(removedTop)
				}
			} else {
				return ""
			}
		}

		for queue.Len() != 0 && x%4 == 0 {
			front := queue.Front().Value.(*element)
			queue.Remove(queue.Front())
			heap.Push(front)
		}

		x++
	}

	return result.String()
}
