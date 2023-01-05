package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newLinkedList() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func TestPrepend(t *testing.T) {

	tests := []struct {
		give []interface{}
		want []interface{}
	}{
		{[]interface{}{2, 3, 4, 5}, []interface{}{2, 3, 4, 5}},
		{[]interface{}{2}, []interface{}{2}},
		{[]interface{}{}, []interface{}{}},
	}

	ll := newLinkedList()

	for i, tt := range tests {
		testname := fmt.Sprintf("TEST:%d", i+1)

		t.Run(testname, func(t *testing.T) {
			for ii, v := range tt.give {
				accval := ll.Prepend(v)
				assert.Equal(t, tt.want[ii], accval)
			}
		})
	}

}

func TestAppend(t *testing.T) {

	tests := []struct {
		give []interface{}
		want []interface{}
	}{
		{[]interface{}{2, 3, 4, 5}, []interface{}{2, 3, 4, 5}},
		{[]interface{}{2}, []interface{}{2}},
		{[]interface{}{}, []interface{}{}},
	}

	ll := newLinkedList()

	for i, tt := range tests {
		testname := fmt.Sprintf("TEST:%d", i+1)

		t.Run(testname, func(t *testing.T) {

			for ii, v := range tt.give {
				accval := ll.Append(v)
				assert.Equal(t, tt.want[ii], accval)
			}

		})
	}

}

func TestSearch(t *testing.T) {

	ll := newLinkedList()

	var inpVals = []interface{}{1, 2, 3, "t", "o", "l", "i", "b"}

	for _, v := range inpVals {
		ll.Append(v)
	}

	tests := []struct {
		give interface{}
		want bool
	}{
		{1, true}, {2, true}, {"o", true}, {"ng", false},
	}

	for i, tt := range tests {

		testname := fmt.Sprintf("TEST:%d", i+1)

		t.Run(testname, func(t *testing.T) {
			accval := ll.Search(tt.give)
			assert.Equal(t, tt.want, accval)
		})
	}
}

func TestDelete(t *testing.T) {

	ll := newLinkedList()

	var inpVals = []interface{}{1, 2, 3, "t", "o", "l", "i", "b"}

	for _, v := range inpVals {
		ll.Append(v)
	}

	tests := []struct {
		give interface{}
		want interface{}
	}{
		{1, 1}, {2, 2}, {"p", nil}, {nil, nil}, {30, nil},
	}

	for i, tt := range tests {

		testname := fmt.Sprintf("TEST:%d", i+1)

		t.Run(testname, func(t *testing.T) {
			accval := ll.Delete(tt.give)
			assert.Equal(t, tt.want, accval)
		})
	}
}
