package main

import (
	"reflect"
	"strconv"
	"testing"

	"strings"
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

		tree := NewTrie()
		for _, s := range c.inserts {
			tree.Insert(s)
		}
		tree.Walk(fn)

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !reflect.DeepEqual(c.edges, result) {
				t.Errorf("\ncase: %v\nexpected: %#v\nresult: %#v", c.inserts, c.edges, result)
			}
		})
	}
}

func TestTrie_Sprint(t *testing.T) {
	data := []string{"test", "toaster", "toasting", "slow", "slowly"}
	expected := []string{
		"t: 27.27% (3)",
		"slow: 18.18% (2)",
		"toast: 18.18% (2)",
		"slowly: 9.09% (1)",
		"test: 9.09% (1)",
		"toaster: 9.09% (1)",
		"toasting: 9.09% (1)",
	}

	cases := []struct {
		count  int
		result string
	}{
		{0, strings.Join(expected, "\n")},
		{len(expected), strings.Join(expected, "\n")},
		{len(expected) + 1, strings.Join(expected, "\n")},
		{1, strings.Join(expected[:1], "\n")},
		{3, strings.Join(expected[:3], "\n")},
	}

	tree := NewTrie()

	result := tree.Sprint(0)
	if result != "" {
		t.Errorf("expected empty string, got %#v", result)
	}

	for _, d := range data {
		tree.Insert(d)
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := tree.Sprint(c.count)
			if result != c.result {
				t.Errorf("expected: %#v\nresult: %#v", c.result, result)
			}
		})
	}
}
