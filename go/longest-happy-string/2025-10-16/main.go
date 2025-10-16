import (
	"container/heap"
	"strings"
)

type maxHeap struct {
	heap []letter
}

func (m maxHeap) Len() int {
	return len(m.heap)
}

func (m maxHeap) Less(i, j int) bool {
	return m.heap[i].frequency > m.heap[j].frequency
}

func (m maxHeap) Swap(i, j int) {
	m.heap[i], m.heap[j] = m.heap[j], m.heap[i]
}

func (m *maxHeap) Push(x any) {
	if x.(letter).frequency > 0 {
		m.heap = append(m.heap, x.(letter))
	}
}

func (m *maxHeap) Pop() any {
	top := m.heap[m.Len()-1]

	m.heap = m.heap[0 : m.Len()-1]

	return top
}

type letter struct {
	frequency int
	value     rune
}

func longestDiverseString(a int, b int, c int) string {
	letters := make([]letter, 0, 3)
	for _, l := range []letter{{a, 'a'}, {b, 'b'}, {c, 'c'}} {
		if l.frequency > 0 {
			letters = append(letters, l)
		}
	}

	h := &maxHeap{
		heap: letters,
	}

	heap.Init(h)

	var result strings.Builder
	var prev letter
	for h.Len() > 0 {
		top := heap.Pop(h).(letter)

		if prev.value == top.value && prev.frequency == 2 {
			if h.Len() == 0 {
				break
			}
			nextTop := heap.Pop(h).(letter)
			prev.value = nextTop.value
			prev.frequency = 1

			nextTop.frequency--
			result.WriteRune(nextTop.value)
			if nextTop.frequency > 0 {
				heap.Push(h, nextTop)
			}

		} else {
			if prev.value != top.value {
				prev.frequency = 0
			}

			prev.value = top.value
			prev.frequency++

			top.frequency--
			result.WriteRune(top.value)
		}

		if top.frequency > 0 {
			heap.Push(h, top)
		}
	}

	return result.String()
}
