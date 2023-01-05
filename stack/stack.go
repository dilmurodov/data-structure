package stack

/**
	*****************
	***** STACK *****
    *****************
**/

type node struct {
	data interface{}
	prev *node
}

type Stack struct {
	top *node
	len int
}

func New() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Push(data interface{}) (val interface{}) {
	nd := &node{data: data, prev: nil}

	nd.prev = s.top
	s.top = nd
	s.len++

	return s.top.data
}

func (s *Stack) Pop() (val interface{}) {

	if s.len == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.len--

	return n.data
}

func (s *Stack) Top() (val interface{}) {
	if s.top == nil {
		return nil
	}
	return s.top.data
}

func (s *Stack) Len() int {
	return s.len
}
