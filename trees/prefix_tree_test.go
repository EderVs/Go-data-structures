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

func equalChains(chain1 []interface{}, chain2 []interface{}) bool {
	for i := range chain1 {
		if chain1[i] != chain2[i] {
			return false
		}
	}
	return true
}

func checkStringValue(t *testing.T, want []string, value string) {
	for _, wantedValue := range want {
		if value == wantedValue {
			return
		}
	}
	t.Errorf("%s is not in %s", value, want)
}

func checkValue(t *testing.T, want [][]interface{}, value []interface{}) {
	for _, wantedValue := range want {
		if equalChains(value, wantedValue) {
			return
		}
	}
	t.Errorf("%s is not in %s", value, want)
}

func checkValues(t *testing.T, want [][]interface{}, content [][]interface{}) {
	if len(want) != len(content) {
		t.Errorf("%s is different than %s", want, content)
	}
	for _, value := range content {
		checkValue(t, want, value)
	}
}

func TestPrefixTree_Insert(t *testing.T) {
	testsInt := []struct {
		name    string
		wantOR  []string
		content [][]interface{}
	}{
		{
			name:    "Slice",
			wantOR:  []string{""},
			content: make([][]interface{}, 0),
		},
		{
			name:    "1 slice",
			wantOR:  []string{"(1(2()))"},
			content: [][]interface{}{{1, 2}},
		},
		{
			name:    "2 slices with same prefix.",
			wantOR:  []string{"(1(2(), 3()))", "(1(2(), 3()))"},
			content: [][]interface{}{{1, 2}, {1, 3}},
		},
	}
	for _, test := range testsInt {
		t.Run(test.name, func(t *testing.T) {
			pT := createPrefixTree(test.content)
			got := pT.String()
			assert.Equal(t, len(test.content), pT.Length())
			checkStringValue(t, test.wantOR, got)
		})
	}
}

func TestPrefixTree_GetValues(t *testing.T) {
	testsInt := []struct {
		name          string
		want, content [][]interface{}
	}{
		{
			name:    "Slice",
			want:    make([][]interface{}, 0),
			content: make([][]interface{}, 0),
		},
		{
			name:    "1 slice",
			want:    [][]interface{}{{1, 2}},
			content: [][]interface{}{{1, 2}},
		},
		{
			name:    "2 slices with same prefix.",
			want:    [][]interface{}{{1, 2}, {1, 3}},
			content: [][]interface{}{{1, 2}, {1, 3}},
		},
		{
			name:    "3 slices with same prefix and final in value in the middle.",
			want:    [][]interface{}{{1, 2}, {1, 3}, {1}},
			content: [][]interface{}{{1, 2}, {1, 3}, {1}},
		},
		{
			name:    "3 slices with same prefix and another slice.",
			want:    [][]interface{}{{1, 2}, {1, 3}, {1}, {2, 4}},
			content: [][]interface{}{{1, 2}, {1, 3}, {1}, {2, 4}},
		},
	}
	for _, test := range testsInt {
		t.Run(test.name, func(t *testing.T) {
			pT := createPrefixTree(test.content)
			got := pT.GetValues()
			assert.Equal(t, len(test.content), pT.Length())
			checkValues(t, test.want, got)
		})
	}
}
