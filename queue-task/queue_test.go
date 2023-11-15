package queue

import "testing"

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	if queue == nil {
		t.Fatalf("queue didnt created")
	}
}

func TestPushQueue(t *testing.T) {
	queue := NewQueue()
	queue.Push(5)
	if queue.sourceList[0] != 5 || queue.size != 1 {
		t.Errorf("push dont work")
	}
}

func TestPopQueue(t *testing.T) {
	queue := NewQueue()
	expected := 5
	queue.Push(expected)
	actual := queue.Pop()
	if queue.size != 0 || expected != actual {
		t.Errorf("pop dont work")
	}
}
