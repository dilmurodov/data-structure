package main

import (
	"github.com/dilmurodov/go-container/heap"
	"github.com/dilmurodov/go-container/linkedlist"
	"github.com/dilmurodov/go-container/queue"
	"github.com/dilmurodov/go-container/stack"
)

func main() {

	h := heap.NewMinHeap()
	h.Len()

	q := queue.New()
	q.Len()

	ll := linkedlist.New()
	ll.Len()

	st := stack.New()
	st.Len()
}
