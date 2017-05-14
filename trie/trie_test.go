package trie_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/31z4/harvest/trie"
)

func TestTrie(t *testing.T) {
	type prefixData struct {
		prefix string
		count  uint
	}

	cases := []struct {
		inserts []string
		edges   []prefixData
	}{
		{
			[]string{},
			[]prefixData{},
		},
		{
			[]string{"test"},
			[]prefixData{
				{"test", 1},
			},
		},
		{
			[]string{"test", "slow"},
			[]prefixData{
				{"test", 1},
				{"slow", 1},
			},
		},
		{
			[]string{"test", "slow", "water", "slower"},
			[]prefixData{
				{"test", 1},
				{"slow", 2},
				{"slower", 1},
				{"water", 1},
			},
		},
		{
			[]string{"tester", "test"},
			[]prefixData{
				{"test", 2},
				{"tester", 1},
			},
		},
		{
			[]string{"test", "team"},
			[]prefixData{
				{"te", 2},
				{"test", 1},
				{"team", 1},
			},
		},
		{
			[]string{"test", "toaster", "toasting", "slow", "slowly"},
			[]prefixData{
				{"t", 3},
				{"test", 1},
				{"toast", 2},
				{"toaster", 1},
				{"toasting", 1},
				{"slow", 2},
				{"slowly", 1},
			},
		},
	}

	for i, c := range cases {
		result := []prefixData{}
		fn := func(prefix string, count uint) {
			result = append(result, prefixData{prefix, count})
		}

		tree := trie.New()
		for _, s := range c.inserts {
			tree.Insert(s)
		}
		tree.Walk(fn)

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !reflect.DeepEqual(c.edges, result) {
				t.Errorf("\ncase: %v\nexpected: %v\nresult: %v", c.inserts, c.edges, result)
			}
		})
	}
}
