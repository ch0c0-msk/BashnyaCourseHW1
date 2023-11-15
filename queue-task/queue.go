package queue

type Type interface{}

type Queue struct {
	sourceList []Type
	size       int
}

func NewQueue() *Queue {
	queue := new(Queue)
	queue.size = 0
	return queue
}

func (q *Queue) Push(val Type) {
	q.sourceList = append([]Type{val}, q.sourceList...)
	q.size++
}

func (q *Queue) Pop() Type {
	if q.size != 0 {
		res := q.sourceList[0]
		q.sourceList = q.sourceList[:(len(q.sourceList) - 1)]
		q.size--
		return res
	} else {
		return nil
	}
}
