package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createIntQueue(content []int) *Queue {
	s := new(Queue)
	for _, value := range content {
		s.Enqueue(value)
	}
	return s
}

func dequeueNthTimes(n int, q *Queue, t *testing.T) *([]int) {
	deleted := make([]int, 0)
	for i := 0; i < n; i++ {
		value, found := q.Dequeue()
		if !found {
			t.Error("Queue must Dequeue.")
		}
		deleted = append(deleted, (*value).(int))
	}
	return &deleted
}

func TestQueue(t *testing.T) {
	testsInt := []struct {
		name, wantedString string
		want               []int
		content            []int
		expectedTail       int
	}{
		{
			name:         "one int array",
			wantedString: "1",
			want:         []int{1},
			content:      []int{1},
			expectedTail: 1,
		},
		{
			name:         "3 int array",
			wantedString: "1 -> 2 -> 3",
			want:         []int{1, 2, 3},
			content:      []int{1, 2, 3},
			expectedTail: 1,
		},
	}
	for _, test := range testsInt {
		t.Run(test.name, func(t *testing.T) {
			q := createIntQueue(test.content)
			got := q.String()
			assert.Equal(t, len(test.content), q.Length())
			assert.Equal(t, test.wantedString, got)
			content := dequeueNthTimes(q.Length(), q, t)
			assert.Equal(t, test.want, *content)
		})
	}
}
