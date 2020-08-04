package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createIntList(content []int) *LinkedList {
	list := new(LinkedList)
	for _, value := range content {
		list.Append(value)
	}
	return list
}

func createIntListWithInsert(content []int) *LinkedList {
	list := new(LinkedList)
	for _, value := range content {
		list.Insert(value)
	}
	return list
}

func popNthTimes(n int, list *LinkedList, t *testing.T) {
	for i := 0; i < n; i++ {
		_, found := list.Pop()
		if !found {
			t.Error("List must pop.")
		}
	}
}

func TestList_Append(t *testing.T) {
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
			name:    "one int array",
			want:    "1",
			content: []int{1},
		},
		{
			name:    "3 int array",
			want:    "1 -> 2 -> 3",
			content: []int{1, 2, 3},
		},
	}
	testsString := []struct {
		name, want string
		content    []string
	}{
		{
			name:    "empty list",
			want:    "",
			content: make([]string, 0),
		},
		{
			name:    "one string array",
			want:    "h",
			content: []string{"h"},
		},
		{
			name:    "3 string array",
			want:    "h -> e -> y",
			content: []string{"h", "e", "y"},
		},
	}
	for _, test := range testsInt {
		t.Run(test.name, func(t *testing.T) {
			list := createIntList(test.content)
			got := list.String()
			assert.Equal(t, len(test.content), list.Length())
			assert.Equal(t, got, test.want)
		})
	}
	for _, test := range testsString {
		t.Run(test.name, func(t *testing.T) {
			list := new(LinkedList)
			for _, value := range test.content {
				list.Append(value)
			}
			got := list.String()
			assert.Equal(t, len(test.content), list.Length())
			assert.Equal(t, test.want, got)
		})
	}
}

func TestList_Insert(t *testing.T) {
	testsInt := []struct {
		name, want   string
		content      []int
		expectedTail int
	}{
		{
			name:         "one int array",
			want:         "1",
			content:      []int{1},
			expectedTail: 1,
		},
		{
			name:         "3 int array",
			want:         "3 -> 2 -> 1",
			content:      []int{1, 2, 3},
			expectedTail: 1,
		},
	}
	for _, test := range testsInt {
		t.Run(test.name, func(t *testing.T) {
			list := createIntListWithInsert(test.content)
			got := list.String()
			assert.Equal(t, len(test.content), list.Length())
			assert.Equal(t, got, test.want)
			assert.Equal(t, test.expectedTail, list.Tail().Value)
		})
	}
}

func TestList_PopAll(t *testing.T) {
	testsPopAll := []struct {
		name, want string
		content    []int
		pops       int
	}{
		{
			name:    "1 pop 1 int array",
			want:    "",
			content: []int{1},
			pops:    1,
		},
		{
			name:    "3 pops 3 int array",
			want:    "",
			content: []int{1, 2, 3},
			pops:    3,
		},
	}
	for _, test := range testsPopAll {
		t.Run(test.name, func(t *testing.T) {
			list := createIntList(test.content)
			popNthTimes(test.pops, list, t)
			got := list.String()
			assert.Equal(t, test.want, got)
		})
	}
	testsTryEmpty := []struct {
		name, want string
		content    []int
		pops       int
	}{
		{
			name:    "Pop in empty list",
			want:    "",
			content: make([]int, 0),
			pops:    1,
		},
	}
	for _, test := range testsTryEmpty {
		t.Run(test.name, func(t *testing.T) {
			list := createIntList(test.content)
			for i := 0; i < test.pops; i++ {
				node, found := list.Pop()
				assert.False(t, found)
				assert.Nil(t, node)
			}
			got := list.String()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestList_PopCheckHeadTail(t *testing.T) {
	testsWithElements := []struct {
		name, want   string
		content      []int
		pops         int
		expectedHead int
		expectedTail int
	}{
		{
			name:         "1 pop 1 int array",
			want:         "1 -> 2 -> 3",
			content:      []int{0, 1, 2, 3},
			pops:         1,
			expectedHead: 1,
			expectedTail: 3,
		},
	}
	for _, test := range testsWithElements {
		t.Run(test.name, func(t *testing.T) {
			list := createIntList(test.content)
			popNthTimes(test.pops, list, t)
			got := list.String()
			assert.Equal(t, got, test.want)
			assert.Equal(t, test.expectedHead, list.Head().Value)
			assert.Equal(t, test.expectedTail, list.Tail().Value)
		})
	}

	testsEmpty := []struct {
		name, want string
		content    []int
	}{
		{
			name:    "4 int array",
			content: []int{0, 1, 2, 3},
		},
	}
	for _, test := range testsEmpty {
		t.Run(test.name, func(t *testing.T) {
			list := createIntList(test.content)
			popNthTimes(list.Length(), list, t)
			got := list.String()
			assert.Equal(t, "", got)
			assert.Equal(t, 0, list.Length())
			assert.Nil(t, list.Head())
			assert.Nil(t, list.Tail())
		})
	}
}

func TestList_PopIth(t *testing.T) {
	testsPopAll := []struct {
		name         string
		want         int
		wantedStr    string
		content      []int
		i            int
		expectedHead int
		expectedTail int
	}{
		{
			name:         "Delete Head",
			want:         1,
			wantedStr:    "2 -> 3",
			content:      []int{1, 2, 3},
			i:            0,
			expectedHead: 2,
			expectedTail: 3,
		},
		{
			name:         "Delete Tail",
			want:         3,
			wantedStr:    "1 -> 2",
			content:      []int{1, 2, 3},
			i:            2,
			expectedHead: 1,
			expectedTail: 2,
		},
		{
			name:         "Delete in the Middle",
			want:         2,
			wantedStr:    "1 -> 3",
			content:      []int{1, 2, 3},
			i:            1,
			expectedHead: 1,
			expectedTail: 3,
		},
	}
	for _, test := range testsPopAll {
		t.Run(test.name, func(t *testing.T) {
			list := createIntList(test.content)
			value, found := list.PopIth(test.i)
			assert.True(t, found)
			assert.Equal(t, test.want, *value)
			got := list.String()
			assert.Equal(t, test.wantedStr, got)
			assert.Equal(t, test.expectedHead, list.Head().Value)
			assert.Equal(t, test.expectedTail, list.Tail().Value)
		})
	}

	testsEmpty := []struct {
		name    string
		content []int
	}{
		{
			name:    "Delete all",
			content: []int{0, 1, 2, 3},
		},
	}
	for _, test := range testsEmpty {
		t.Run(test.name, func(t *testing.T) {
			list := createIntList(test.content)
			n := list.Length()
			for i := 0; i < n; i++ {
				_, found := list.PopIth(0)
				if !found {
					t.Error("List must pop.")
				}
			}
			got := list.String()
			assert.Equal(t, "", got)
			assert.Equal(t, 0, list.Length())
			assert.Nil(t, list.Head())
			assert.Nil(t, list.Tail())
		})
	}
}

func TestList_GetIthNode(t *testing.T) {
	testsInList := []struct {
		name    string
		want    int
		content []int
		i       int
	}{
		{
			name:    "1st",
			want:    2,
			content: []int{1, 2, 3},
			i:       1,
		},
		{
			name:    "0th",
			want:    1,
			content: []int{1},
			i:       0,
		},
		{
			name:    "2nd",
			want:    3,
			content: []int{1, 2, 3},
			i:       2,
		},
	}
	for _, test := range testsInList {
		t.Run(test.name, func(t *testing.T) {
			list := createIntList(test.content)
			node, found := list.GetIthNode(test.i)
			if !found {
				t.Error("Element must exist.")
			}
			assert.Equal(t, node.Value, test.want)
		})
	}

	testsOutList := []struct {
		name    string
		content []int
		i       int
	}{
		{
			name:    "Negative value",
			content: []int{1, 2, 3},
			i:       -1,
		},
		{
			name:    "Empty list",
			content: make([]int, 0),
			i:       0,
		},
		{
			name:    "Lenght of list",
			content: []int{1, 2, 3},
			i:       3,
		},
	}
	for _, test := range testsOutList {
		t.Run(test.name, func(t *testing.T) {
			list := createIntList(test.content)
			_, found := list.GetIthNode(test.i)
			assert.False(t, found)
		})
	}
}
