package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createPrefixTree(content [][]interface{}) *PrefixTree {
	pT := new(PrefixTree)
	for _, value := range content {
		pT.Insert(value)
	}
	return pT
}

func TestPrefixTree_Insert(t *testing.T) {
	testsInt := []struct {
		name, want string
		content    [][]interface{}
	}{
		{
			name:    "Slice",
			want:    "",
			content: make([][]interface{}, 0),
		},
		{
			name:    "1 slice",
			want:    "(1(2()))",
			content: [][]interface{}{{1, 2}},
		},
		{
			name:    "2 slices with same prefix.",
			want:    "(1(2(), 3()))",
			content: [][]interface{}{{1, 2}, {1, 3}},
		},
	}
	for _, test := range testsInt {
		t.Run(test.name, func(t *testing.T) {
			pT := createPrefixTree(test.content)
			got := pT.String()
			assert.Equal(t, len(test.content), pT.Length())
			assert.Equal(t, got, test.want)
		})
	}
}
