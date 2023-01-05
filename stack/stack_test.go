package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newStack() *Stack {
	return &Stack{nil, 0}
}

func TestPush(t *testing.T) {
	var tests = []struct {
		give []interface{}
		want []interface{}
	}{
		{[]interface{}{3}, []interface{}{3}},
		{[]interface{}{3, 4, 5}, []interface{}{3, 4, 5}},
		{[]interface{}{'t', 'm', 'k'}, []interface{}{'t', 'm', 'k'}},
		{[]interface{}{"year 2023"}, []interface{}{"year 2023"}},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("TEST:%d", i+1)

		s := newStack()

		t.Run(testname, func(t *testing.T) {

			for ii, v := range tt.give {
				pd := s.Push(v)

				assert.Equal(t, tt.want[ii], pd)

			}
		})
	}
}

func TestPop(t *testing.T) {

	var tests = []struct {
		give []interface{}
		want interface{}
	}{
		{give: []interface{}{2, 3, 5}, want: 5},
		{give: []interface{}{'t', 'o'}, want: 'o'},
		{give: []interface{}{"2023", "2024"}, want: "2024"},
		{give: []interface{}{}, want: nil},
		{give: []interface{}{1}, want: 1},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("TEST:%d", i+1)

		s := newStack()

		t.Run(testname, func(t *testing.T) {

			for _, v := range tt.give {
				s.Push(v)
			}

			pd := s.Pop()
			assert.Equal(t, tt.want, pd)
		})
	}
}
