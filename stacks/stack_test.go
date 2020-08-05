package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createIntStack(content []int) *Stack {
	s := new(Stack)
	for _, value := range content {
		s.Add(value)
	}
	return s
}

func popNthTimes(n int, s *Stack, t *testing.T) *([]int) {
	deleted := make([]int, 0)
	for i := 0; i < n; i++ {
		value, found := s.Pop()
		if !found {
			t.Error("Stack must pop.")
		}
		deleted = append(deleted, (*value).(int))
	}
	return &deleted
}

func TestStack(t *testing.T) {
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
			wantedString: "3 -> 2 -> 1",
			want:         []int{3, 2, 1},
			content:      []int{1, 2, 3},
			expectedTail: 1,
		},
	}
	for _, test := range testsInt {
		t.Run(test.name, func(t *testing.T) {
			s := createIntStack(test.content)
			got := s.String()
			assert.Equal(t, len(test.content), s.Length())
			assert.Equal(t, got, test.wantedString)
			content := popNthTimes(s.Length(), s, t)
			assert.Equal(t, test.want, *content)
		})
	}
}
