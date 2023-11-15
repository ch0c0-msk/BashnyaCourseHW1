package deque

type Type interface{}

type Node struct {
	prev  *Node
	next  *Node
	value Type
}

type Deque struct {
	first *Node
	last  *Node
	size  int
}

func NewDeque() *Deque {
	deque := new(Deque)
	deque.size = 0
	return deque
}

func (d *Deque) PushFront(value Type) {
	node := new(Node)
	node.value = value
	if d.size == 0 {
		d.first = node
		d.last = node
	} else {
		d.first.prev = node
		node.next = d.first
		d.first = node
	}
	d.size++
}

func (d *Deque) PushBack(value Type) {
	node := new(Node)
	node.value = value
	if d.size == 0 {
		d.first = node
		d.last = node
	} else {
		d.last.next = node
		node.prev = d.last
		d.last = node
	}
	d.size++
}

func (d *Deque) PopFront() Type {
	if d.size != 0 {
		res := d.first.value
		if d.size == 1 {
			d.first = nil
			d.last = nil
		} else {
			d.first = d.first.next
			d.first.prev = nil
		}
		d.size--
		return res
	} else {
		return nil
	}
}

func (d *Deque) PopBack() Type {
	if d.size != 0 {
		res := d.last.value
		if d.size == 1 {
			d.last = nil
			d.first = nil
		} else {
			d.last = d.last.prev
			d.last.next = nil
		}
		d.size--
		return res
	} else {
		return nil
	}
}
