package redblack

import (
	"reflect"
	"testing"
)

func cmpInt(a, b int) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func TestDeleteReplacesNodeWithSuccessorData(t *testing.T) {
	tree := New[int, int](cmpInt, 0)

	for _, v := range []int{20, 10, 30, 25, 40} {
		if !tree.Insert(v, v*10) {
			t.Fatalf("Insert(%d) failed", v)
		}
	}

	key, value, ok := tree.Delete(20)
	if !ok || key != 20 || value != 200 {
		t.Fatalf("Delete(20) = (%v, %v, %v), want (20, 200, true)", key, value, ok)
	}

	if gotValue, ok := tree.Get(20); ok {
		t.Fatalf("Get(20) = (%v, %v), want (_, false)", gotValue, ok)
	}

	for _, v := range []int{10, 25, 30, 40} {
		gotValue, ok := tree.Get(v)
		if !ok || gotValue != v*10 {
			t.Fatalf("Get(%d) = (%v, %v), want (%d, true)", v, gotValue, ok, v*10)
		}
	}

	if got := tree.Len(); got != 4 {
		t.Fatalf("Len() = %d, want 4", got)
	}
}

func TestDeleteSequenceDoesNotPanicOnSentinelSibling(t *testing.T) {
	tree := New[int, int](cmpInt, 0)

	for _, v := range []int{2, 1, 3, 4} {
		if !tree.Insert(v, v) {
			t.Fatalf("Insert(%d) failed", v)
		}
	}

	for _, v := range []int{1, 2, 3, 4} {
		key, value, ok := tree.Delete(v)
		if !ok || key != v || value != v {
			t.Fatalf("Delete(%d) = (%v, %v, %v), want (%d, %d, true)", v, key, value, ok, v, v)
		}
	}

	if got := tree.Len(); got != 0 {
		t.Fatalf("Len() = %d, want 0", got)
	}
}

func TestMinMaxAndRemoveMax(t *testing.T) {
	tree := New[int, string](cmpInt, 0)

	if _, _, ok := tree.Min(); ok {
		t.Fatal("Min() on empty tree returned ok=true")
	}
	if _, _, ok := tree.Max(); ok {
		t.Fatal("Max() on empty tree returned ok=true")
	}
	if _, _, ok := tree.RemoveMax(); ok {
		t.Fatal("RemoveMax() on empty tree returned ok=true")
	}

	for _, item := range []struct {
		key   int
		value string
	}{
		{20, "twenty"},
		{10, "ten"},
		{30, "thirty"},
	} {
		if !tree.Insert(item.key, item.value) {
			t.Fatalf("Insert(%d) failed", item.key)
		}
	}

	minKey, minValue, ok := tree.Min()
	if !ok || minKey != 10 || minValue != "ten" {
		t.Fatalf("Min() = (%v, %v, %v), want (10, ten, true)", minKey, minValue, ok)
	}

	maxKey, maxValue, ok := tree.Max()
	if !ok || maxKey != 30 || maxValue != "thirty" {
		t.Fatalf("Max() = (%v, %v, %v), want (30, thirty, true)", maxKey, maxValue, ok)
	}

	removedKey, removedValue, ok := tree.RemoveMax()
	if !ok || removedKey != 30 || removedValue != "thirty" {
		t.Fatalf("RemoveMax() = (%v, %v, %v), want (30, thirty, true)", removedKey, removedValue, ok)
	}
}

func TestWalkOrdersAndEarlyStop(t *testing.T) {
	tree := New[int, int](cmpInt, 0)
	for _, v := range []int{20, 10, 30, 5, 15, 25, 35} {
		tree.Insert(v, v)
	}

	var inOrder []int
	tree.Walk(InOrder, func(key, value int) bool {
		inOrder = append(inOrder, key)
		return true
	})
	if want := []int{5, 10, 15, 20, 25, 30, 35}; !reflect.DeepEqual(inOrder, want) {
		t.Fatalf("InOrder walk = %v, want %v", inOrder, want)
	}

	var preOrder []int
	tree.Walk(PreOrder, func(key, value int) bool {
		preOrder = append(preOrder, key)
		return true
	})
	if want := []int{20, 10, 5, 15, 30, 25, 35}; !reflect.DeepEqual(preOrder, want) {
		t.Fatalf("PreOrder walk = %v, want %v", preOrder, want)
	}

	var limited []int
	tree.Walk(InOrder, func(key, value int) bool {
		limited = append(limited, key)
		return len(limited) < 3
	})
	if want := []int{5, 10, 15}; !reflect.DeepEqual(limited, want) {
		t.Fatalf("early-stop walk = %v, want %v", limited, want)
	}
}
