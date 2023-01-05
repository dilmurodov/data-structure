package heap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newHeap() HeapI {
	return NewMinHeap()
}

func TestAdd(t *testing.T) {

	tests := []struct {
		give []int64
		want []int
	}{
		{[]int64{4, 6, 3}, []int{1, 2, 3}},
		{[]int64{4, 6, 3, 5, 7, 8}, []int{1, 2, 3, 4, 5, 6}},
		{[]int64{1}, []int{1}},
		{[]int64{}, []int{0}},
	}

	for i, tt := range tests {
		h := newHeap()

		testname := fmt.Sprintf("TEST:%d", i+1)

		t.Run(testname, func(t *testing.T) {

			for ii, v := range tt.give {
				accval := h.Add(v)
				assert.Equal(t, tt.want[ii], accval)
			}
		})
	}
}

func TestPick(t *testing.T) {

	tests := []struct {
		give []int64
		want int64
	}{
		{[]int64{4, 6, 3}, 3},
		{[]int64{4, 6, 3, 5, 7, 8}, 3},
		{[]int64{1}, 1},
		{[]int64{}, 0},
	}

	for i, tt := range tests {

		h := newHeap()

		testname := fmt.Sprintf("TEST:%d", i+1)

		for _, v := range tt.give {
			h.Add(v)
		}

		t.Run(testname, func(t *testing.T) {
			accval, _ := h.Pick()
			assert.Equal(t, tt.want, accval)
		})
	}

}

func TestPoll(t *testing.T) {

	inputVals := []int64{4, 6, 3, 5, 7, 8}

	h := newHeap()

	for _, v := range inputVals {
		h.Add(v)
	}

	tests := []struct {
		want struct {
			val int64
			ok  bool
		}
	}{
		{
			want: struct {
				val int64
				ok  bool
			}{3, true},
		},
		{
			want: struct {
				val int64
				ok  bool
			}{4, true},
		},
		{
			want: struct {
				val int64
				ok  bool
			}{5, true},
		},
		{
			want: struct {
				val int64
				ok  bool
			}{6, true},
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("TEST:%d", i+1)

		t.Run(testname, func(t *testing.T) {

			val, ok := h.Poll()

			assert.Equal(t, tt.want, struct {
				val int64
				ok  bool
			}{val, ok})

		})
	}
}

func TestLen(t *testing.T) {

	tests := []struct {
		give []int64
		want int
	}{
		{[]int64{1, 2, 3, 4, 5}, 5},
		{[]int64{1, 2}, 2},
		{[]int64{1}, 1},
		{[]int64{}, 0},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("TEST:%d", i+1)

		t.Run(testname, func(t *testing.T) {

			h := newHeap()

			for _, v := range tt.give {
				h.Add(v)
			}

			accval := h.Len()

			assert.Equal(t, tt.want, accval)

		})
	}
}
