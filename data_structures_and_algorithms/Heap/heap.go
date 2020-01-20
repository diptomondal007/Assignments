package heap

//Heap type
type Heap struct {
	Items []int
}

//GetLeftIndex : Get left index of a heaf node
func (h *Heap) GetLeftIndex(index int) int {
	return index*2 + 1
}

//GetRightIndex : Get Right index of a heaf node
func (h *Heap) GetRightIndex(index int) int {
	return index*2 + 2
}

//GetParentIndex : To get the parent index of a node
func (h *Heap) GetParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

//HasRight : To check if a node has a element in right
func (h *Heap) HasRight(index int) bool {
	return h.GetRightIndex(index) < len(h.Items)
}

//HasLeft : To check if a node has a left index
func (h *Heap) HasLeft(index int) bool {
	return h.GetRightIndex(index) < len(h.Items)
}

//HasParent : To check if a node has a parent
func (h *Heap) HasParent(index int) bool {
	return h.GetParentIndex(index) >= 0
}

//Left : Get the left element of a node
func (h *Heap) Left(index int) int {
	return h.Items[h.GetLeftIndex(index)]
}

//Right : Get the right element of a node
func (h *Heap) Right(index int) int {
	return h.Items[h.GetRightIndex(index)]
}

//Parent : Get the parent node element
func (h *Heap) Parent(index int) int {
	return h.Items[h.GetParentIndex(index)]
}

// Swap : swap values of two nodes at specified indeces
func (h *Heap) Swap(indexOne, indexTwo int) {
	h.Items[indexOne], h.Items[indexTwo] = h.Items[indexTwo], h.Items[indexOne]
}
