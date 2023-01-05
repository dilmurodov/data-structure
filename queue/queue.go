/**
	*****************
	***** QUEUE *****
    *****************
**/

package queue

type node struct {
	next *node
	data interface{}
}

type Queue struct {
	head *node
	tail *node
	len  int
}

func New() *Queue {
	return &Queue{
		nil, nil, 0,
	}
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) Push(data interface{}) (val interface{}) {

	nd := &node{next: nil, data: data}

	if q.len == 0 {
		q.head = nd
		q.tail = nd
	} else {
		q.tail.next = nd
		q.tail = nd
	}

	q.len++

	return q.tail.data
}

func (q *Queue) Pop() (val interface{}) {

	if q.len == 0 {
		return nil
	}

	if q.len == 1 {
		val = q.head.data
		q.head = nil
		q.tail = nil
	} else {
		val = q.head.data
		q.head = q.head.next
	}

	return val
}

func (q *Queue) Top() (val interface{}) {
	if q.head == nil {
		return nil
	}
	return q.head.data
}
