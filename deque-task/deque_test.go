package deque

import "testing"

func TestNewDeque(t *testing.T) {

	deque := NewDeque()
	if deque == nil {
		t.Fatalf("deque didnt created")
	}
}

func TestPushFront(t *testing.T) {
	deque := NewDeque()
	expected := "some_value"
	deque.PushFront(expected)
	actual := deque.first.value
	if deque.size != 1 || expected != actual {
		t.Errorf("push front dont work")
	}
}

func TestPushBack(t *testing.T) {
	deque := NewDeque()
	expected := "some_value"
	deque.PushBack(expected)
	actual := deque.last.value
	if deque.size != 1 || expected != actual {
		t.Errorf("push back dont work")
	}
}

func TestPopFront(t *testing.T) {
	deque := NewDeque()
	expected := 5
	deque.PushFront(expected)
	actual := deque.PopFront()
	if deque.size != 0 || expected != actual {
		t.Errorf("pop front dont work")
	}
}

func TestPopBack(t *testing.T) {
	deque := NewDeque()
	expected := 5
	deque.PushBack(expected)
	actual := deque.PopBack()
	if deque.size != 0 || expected != actual {
		t.Errorf("pop front dont work")
	}
}
