import (
	"container/heap"
	"math"
)

type minHeap []point

type point = []int

func distanceFromOrigin(p point) float64 {
	return math.Sqrt(float64(p[0]*p[0] + p[1]*p[1]))
}

func (mh minHeap) Len() int {
	return len(mh)
}

func (mh minHeap) Less(i, j int) bool {
	return distanceFromOrigin(mh[i]) < distanceFromOrigin(mh[j])
}

func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minHeap) Push(x any) {
	*mh = append(*mh, x.(point))
}

func (mh *minHeap) Pop() any {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	result := make([][]int, 0, k)
	h := &minHeap{}
	heap.Init(h)

	for _, point := range points {
		heap.Push(h, point)
	}

	for range k {
		result = append(result, heap.Pop(h).(point))
	}

	return result
}
