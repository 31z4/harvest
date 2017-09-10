package main

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &Heap{
		&KeySize{"test1", 1},
		&KeySize{"test2", 1},
		&KeySize{"test3", 3},
	}

	heap.Init(h)
	heap.Push(h, &KeySize{"test4", 2})

	expected := []*KeySize{
		{"test3", 3},
		{"test4", 2},
		{"test1", 1},
		{"test2", 1},
	}
	result := []*KeySize{}

	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(*KeySize))
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected: %#v\nresult: %#v", expected, result)
	}
}
