package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newQueue() *Queue {
	q := New()
	return q
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

		q := newQueue()

		t.Run(testname, func(t *testing.T) {

			for ii, v := range tt.give {
				pd := q.Push(v)

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
		{give: []interface{}{2, 3, 5}, want: 2},
		{give: []interface{}{'t', 'o'}, want: 't'},
		{give: []interface{}{"2023", "2024"}, want: "2023"},
		{give: []interface{}{}, want: nil},
		{give: []interface{}{1}, want: 1},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("TEST:%d", i+1)

		q := newQueue()

		t.Run(testname, func(t *testing.T) {

			for _, v := range tt.give {
				q.Push(v)
			}

			pd := q.Pop()
			assert.Equal(t, tt.want, pd)
		})
	}
}
