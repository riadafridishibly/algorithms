package heap

// An implementation of max heap
// the `less` function is hard coded just for simplicity

type MaxHeap struct {
	data []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (m *MaxHeap) Add(e int) {
	m.data = append(m.data, e)
	up(m.data, len(m.data)-1)
}

func (m *MaxHeap) Len() int {
	return len(m.data)
}

func (m *MaxHeap) PopMax() int {
	e := m.data[0]
	n := len(m.data)
	// swap with the last element
	m.data[0], m.data[n-1] = m.data[n-1], m.data[0]
	// shrink the array
	m.data = m.data[:n-1]
	// preserve the heap property
	down(m.data, 0, len(m.data))
	return e
}

func leftChild(i int) int {
	return i*2 + 1
}

func rightChild(i int) int {
	return i*2 + 2
}

func parent(i int) int {
	// -1 / 2 == 0
	return (i - 1) / 2
}

// Heapify runs heapify algorithm on array
func Heapify(array []int) {
	n := len(array)
	for i := len(array) - 1; i >= 0; i-- {
		down(array, i, n)
	}
}

func HeapSort(array []int) {
	// 1. Create heap (this way the maximum element will be placed at index 0)
	Heapify(array)

	for n := len(array) - 1; n > 0; n-- {
		// 2. Remove the max element (swap it with the last element)
		array[0], array[n] = array[n], array[0]
		// 3. Preserve the heap property for the rest of the elements
		down(array, 0, n)
	}
}

// i = parent element
// n = length of array / last+1 index where down should stop
func down(array []int, i, n int) {
	for {
		l, r := leftChild(i), rightChild(i)
		if l >= n && r >= n { // overflow
			return
		}

		// take the largest child
		j := l
		if r < n && array[r] > array[l] {
			j = r
		}

		// parent is already larger then the two children
		if array[i] > array[j] {
			return
		}

		array[i], array[j] = array[j], array[i]
		i = j
	}
}

func up(array []int, i int) {
	for {
		p := parent(i)
		if p == i /* parent(0) == 0 */ || array[p] > array[i] /* parent already larger */ {
			break
		}
		array[p], array[i] = array[i], array[p]
		i = p
	}
}
