package heap

import (
	"fmt"
)

type Heap struct {
	heap []int
}

func (h *Heap) Add(val int) {
	h.heap = append(h.heap, val)
	h.heapify()
}

func (h *Heap) heapify() {
	index := len(h.heap) - 1
	for index > 0 && h.heap[index] < h.heap[(index-1)/2] {
		h.heap[index], h.heap[(index-1)/2] = h.heap[(index-1)/2], h.heap[index]
		index = (index - 1) / 2
	}
}

func (h *Heap) Delete(val int) bool {
	index := indexOf(h.heap, val)
	if index < 0 {
		return false
	}

	return false
}

func indexOf(heap []int, val int) int {
	for i := 0; i < len(heap); i++ {
		if heap[i] == val {
			return i
		}
	}
	return -1
}

func (h *Heap) Contains(val int) bool {
	return indexOf(h.heap, val) != -1
}

func (h *Heap) String() {
	fmt.Println(h.heap)
	fmt.Println("the heap length:", len(h.heap))
}
