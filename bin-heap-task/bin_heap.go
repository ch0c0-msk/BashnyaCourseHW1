package bin_heap

import "errors"

type BinHeap struct {
	data []float64
	size int
}

func NewBinHeap(sourceList []float64) *BinHeap {
	binHeap := new(BinHeap)
	binHeap.data = sourceList
	binHeap.size = len(binHeap.data)
	binHeap.build()
	return binHeap
}

func (b *BinHeap) build() {
	for i := (b.size/2 - 1); i >= 0; i-- {
		b.siftDown(i)
	}
}

func (b *BinHeap) siftDown(i int) {
	var leftChild, rightChild, largestChild int
	for {

		leftChild = 2*i + 1
		rightChild = 2*i + 2
		largestChild = i

		if leftChild < b.size && b.data[leftChild] > b.data[largestChild] {
			largestChild = leftChild
		}

		if rightChild < b.size && b.data[rightChild] > b.data[largestChild] {
			largestChild = rightChild
		}

		if largestChild == i {
			break
		}

		temp := b.data[i]
		b.data[i] = b.data[largestChild]
		b.data[largestChild] = temp
		i = largestChild
	}
}

func (b *BinHeap) siftUp(i int) {
	var parent int
	for {
		if i == 0 {
			break
		}
		if i%2 == 0 {
			parent = i/2 - 1
		} else {
			parent = i / 2
		}
		if b.data[i] > b.data[parent] {
			b.data[i], b.data[parent] = b.data[parent], b.data[i]
			i = parent
		} else {
			break
		}
	}
}

func (b *BinHeap) Top() (float64, error) {
	if b.size == 0 {
		return 0, errors.New("the heap is empty")
	} else {
		return b.data[0], nil
	}
}

func (b *BinHeap) Insert(elem float64) {
	b.data = append(b.data, elem)
	elemIndex := len(b.data) - 1
	b.siftUp(elemIndex)
	b.size = len(b.data)
}

func (b *BinHeap) ExtractTop() (float64, error) {
	if b.size == 0 {
		return 0, errors.New("the heap is empty")
	} else {
		top := b.data[0]
		b.data = append(b.data[1:len(b.data)-2], b.data[len(b.data)-1])
		b.siftDown(0)
		b.size = len(b.data)
		return top, nil
	}
}
