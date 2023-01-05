/** ********************
	***** MIN HEAP *****
    ********************
          10
		/    \
	   13     23
	 /   \   /
	17   20 25
**/

package heap

type MinHeap struct {
	data []int64
	len  int
}

type HeapI interface {
	Add(data int64) (len int)
	Poll() (val int64, ok bool)
	Show() (data []int64)
	Pick() (val int64, ok bool)
	Len() int
}

func NewMinHeap() HeapI {
	return &MinHeap{len: 0, data: nil}
}

func (h *MinHeap) Poll() (val int64, ok bool) {

	if h.len == 0 {
		return 0, false
	}

	var root = h.data[0]
	h.data = h.data[1:]

	h.len--
	h.heapifyDown()

	return root, true
}

func (h *MinHeap) Add(data int64) (len int) {

	h.data = append(h.data, data)
	
	h.len++

	h.heapifyUp()

	return h.len
}

func (h *MinHeap) Show() (data []int64) {
	return h.data
}

func (h *MinHeap) Len() int {
	return h.len
}

func (h *MinHeap) Pick() (val int64, ok bool) {
	if h.len == 0 {
		ok = false
		return
	}
	return h.data[0], true
}

func (h *MinHeap) swap(index1 int, index2 int) {
	h.data[index1], h.data[index2] = h.data[index2], h.data[index1]
}

func (h *MinHeap) heapifyUp() {
	var index = h.len - 1

	for h.hasParent(index) && h.parent(index) > h.data[index] {
		h.swap(index, h.getParentIndex(index))
		index = h.getParentIndex(index)
	}
}

func (h *MinHeap) heapifyDown() {
	var index = 0

	for h.hasLeftChild(index) {
		smallChildIndex := h.getLeftChildIndex(index)

		if h.hasRightChild(index) && h.rightChild(index) < h.leftChild(index) {
			smallChildIndex = h.getLeftChildIndex(index)
		}

		if h.data[index] < h.data[smallChildIndex] {
			break
		} else {
			h.swap(index, smallChildIndex)
		}

		index = smallChildIndex

	}
}

// get index of node

func (h *MinHeap) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *MinHeap) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *MinHeap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

// check node is exists

func (h *MinHeap) hasLeftChild(index int) bool {
	return h.getLeftChildIndex(index) < h.len
}

func (h *MinHeap) hasRightChild(index int) bool {
	return h.getRightChildIndex(index) < h.len
}

func (h *MinHeap) hasParent(index int) bool {
	return h.getParentIndex(index) >= 0
}

// get node

func (h *MinHeap) leftChild(index int) int64 {
	return h.data[h.getLeftChildIndex(index)]
}

func (h *MinHeap) rightChild(index int) int64 {
	return h.data[h.getRightChildIndex(index)]
}

func (h *MinHeap) parent(index int) int64 {
	return h.data[h.getParentIndex(index)]
}
