/**
	***********************
	***** LINKED LIST *****
    ***********************
**/

package linkedlist

type node struct {
	next *node
	data interface{}
}

type LinkedList struct {
	head *node
	tail *node
	len  int
}

func New() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func (ll *LinkedList) Prepend(data interface{}) (val interface{}) {

	nd := &node{next: nil, data: data}

	if ll.head == nil {
		ll.head = nd
		ll.tail = nd
	} else {
		nd.next = ll.head
		ll.head = nd
	}

	ll.len++

	return ll.head.data
}

func (ll *LinkedList) Append(data interface{}) (val interface{}) {
	nd := &node{data: data, next: nil}

	if ll.head == nil {
		ll.head = nd
		ll.tail = nd
	} else {
		ll.tail.next = nd
		ll.tail = nd
	}

	ll.len++

	return ll.tail.data
}

func (ll *LinkedList) Search(data interface{}) bool {

	for curr := ll.head; curr != nil; {
		if curr.data == data {
			return true
		} else {
			curr = curr.next
		}
	}

	return false
}

func (ll *LinkedList) Delete(data interface{}) (val interface{}) {

	if ll.head == nil {
		return nil
	} else if ll.head.data == data {
		ll.head = ll.head.next
		return data
	} else {
		tHead := ll.head
		for curr := tHead; curr.next != nil; {
			if curr.next.data == data {
				curr.next = curr.next.next
				ll.head = tHead
				ll.len--
				return data
			} else {
				curr = curr.next
			}
		}
	}

	return nil
}

func (ll *LinkedList) Len() int {
	return ll.len
}
