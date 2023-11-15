package bin_heap

import "testing"

func findMax(list []float64) float64 {
	max := list[0]
	for _, v := range list {
		if v > max {
			max = v
		}
	}
	return max
}
func TestNewBinHeap(t *testing.T) {

	testSlice := []float64{1, 3, 5, 6, 9, 7, 8, 8, 11, 15}
	binHeap := NewBinHeap(testSlice)
	if binHeap == nil {
		t.Fatalf("binary heap didnt created")
	}

	if binHeap.size != len(testSlice) || binHeap.data[0] != findMax(testSlice) {
		t.Errorf("binary heap is not correct")
	}
}

func TestTop(t *testing.T) {

	testSlice := []float64{1, 3, 5, 6, 9, 7, 8, 8, 11, 15}
	binHeap := NewBinHeap(testSlice)
	expected := findMax(testSlice)
	actual, err := binHeap.Top()
	if err != nil || expected != actual {
		t.Errorf("top method dont work")
	}
}

func TestInsert(t *testing.T) {

	testSlice := []float64{1, 3, 5, 6, 9, 7, 8, 8, 11, 15}
	binHeap := NewBinHeap(testSlice)
	expected := float64(13)
	binHeap.Insert(expected)
	actual := binHeap.data[1]
	if expected != actual {
		t.Errorf("insert method dont work")
	}
}

func TestExtractTopCorrectReturn(t *testing.T) {
	testSlice := []float64{1, 3, 5, 6, 9, 7, 8, 8, 11, 15}
	binHeap := NewBinHeap(testSlice)
	expected := float64(15)
	actual, err := binHeap.ExtractTop()
	if expected != actual || err != nil {
		t.Errorf("extract top method dont work")
	}
}

func TestExtractTopCorrectHeapAfterExtract(t *testing.T) {
	testSlice := []float64{1, 3, 5, 6, 9, 7, 8, 8, 11, 15}
	binHeap := NewBinHeap(testSlice)
	_, err := binHeap.ExtractTop()
	expected := float64(11)
	actual := binHeap.data[0]
	if expected != actual || err != nil {
		t.Errorf("extract top method dont work")
	}
}
