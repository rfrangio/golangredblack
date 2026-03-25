package redblack

import "testing"

func cmpInt(key1 interface{}, key2 interface{}) int {
	intKey1 := key1.(int)
	intKey2 := key2.(int)

	if intKey1 == intKey2 {
		return 0
	}
	if intKey1 > intKey2 {
		return 1
	}
	return -1
}

func TestDeleteReplacesNodeWithSuccessorData(t *testing.T) {
	tree := ConstructRedBlackTree(cmpInt, 0)

	for _, v := range []int{20, 10, 30, 25, 40} {
		if tree.Insert(v, v*10) != 1 {
			t.Fatalf("insert(%d) failed", v)
		}
	}

	key, value := tree.Delete(20)
	if key != 20 || value != 200 {
		t.Fatalf("Delete(20) = (%v, %v), want (20, 200)", key, value)
	}

	if gotKey, gotValue := tree.Search(20); gotKey != nil || gotValue != nil {
		t.Fatalf("Search(20) = (%v, %v), want (nil, nil)", gotKey, gotValue)
	}

	for _, v := range []int{10, 25, 30, 40} {
		gotKey, gotValue := tree.Search(v)
		if gotKey != v || gotValue != v*10 {
			t.Fatalf("Search(%d) = (%v, %v), want (%d, %d)", v, gotKey, gotValue, v, v*10)
		}
	}

	if count := tree.DoNodeCount(); count != 4 {
		t.Fatalf("DoNodeCount() = %d, want 4", count)
	}
}

func TestDeleteSequenceDoesNotPanicOnSentinelSibling(t *testing.T) {
	tree := ConstructRedBlackTree(cmpInt, 0)

	for _, v := range []int{2, 1, 3, 4} {
		if tree.Insert(v, v) != 1 {
			t.Fatalf("insert(%d) failed", v)
		}
	}

	for _, v := range []int{1, 2, 3, 4} {
		key, value := tree.Delete(v)
		if key != v || value != v {
			t.Fatalf("Delete(%d) = (%v, %v), want (%d, %d)", v, key, value, v, v)
		}
	}

	if count := tree.DoNodeCount(); count != 0 {
		t.Fatalf("DoNodeCount() = %d, want 0", count)
	}
}
