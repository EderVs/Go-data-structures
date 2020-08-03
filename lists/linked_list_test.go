package lists

import "testing"

func TestList_InsertInt(t *testing.T) {
	testsInt := []struct {
		name, want string
		content    []int
	}{
		{
			name:    "empty list",
			want:    "",
			content: make([]int, 0),
		},
		{
			name:    "one element array",
			want:    "1",
			content: []int{1},
		},
		{
			name:    "2 element array",
			want:    "1 -> 2 -> 3",
			content: []int{1, 2, 3},
		},
	}
	for _, test := range testsInt {
		t.Run(test.name, func(t *testing.T) {
			list := new(LinkedList)
			for _, value := range test.content {
				list.Insert(value)
			}
			got := list.String()
			if test.want != got {
				t.Errorf("String() = %s, want %s", got, test.want)
			}
		})
	}
}
