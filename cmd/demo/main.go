package main

import (
	"fmt"
	"golangredblack/redblack"
	"math/rand"
	"time"
)

func cmpString(key1 string, key2 string) int {
	if key1 == key2 {
		return 0
	}
	if key1 > key2 {
		return 1
	}
	return -1
}

func cmpInt(key1 int, key2 int) int {
	if key1 == key2 {
		return 0
	}
	if key1 > key2 {
		return 1
	}
	return -1
}

func main() {
	rbtree := redblack.New[int, int](cmpInt, 0)
	rbtreec := redblack.New[string, string](cmpString, 1000)
	var numObjects int = 1939347
	var dups int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rbtreec.Insert("a", "a")

	fmt.Printf("Inserting %d sequential objects \n", numObjects)

	for i := 0; i < numObjects; i++ {
		rbtree.Insert(i, i)
	}

	nodeCount := rbtree.Len()
	fmt.Printf("node count %d \n", nodeCount)
	ldepth, rdepth := rbtree.SubtreeDepths()
	fmt.Printf("depths %d, %d \n", ldepth, rdepth)

	for i := 0; i < numObjects; i++ {
		rbtree.RemoveMax()
	}

	nodeCount = rbtree.Len()
	fmt.Printf("node count %d \n", nodeCount)

	fmt.Printf("Inserting %d random objects \n", numObjects)
	nodeCount, ldepth, rdepth = 0, 0, 0

	var entry int
	for i := 0; i < numObjects; i++ {
		entry = r.Int()
		if !rbtree.Insert(entry, entry) {
			dups++
		}
	}

	nodeCount = rbtree.Len()
	fmt.Printf("node count %d \n", nodeCount+dups)
	ldepth, rdepth = rbtree.SubtreeDepths()
	fmt.Printf("depths %d, %d \n", ldepth, rdepth)

	for i := 0; i < numObjects; i++ {
		rbtree.RemoveMax()
	}

	nodeCount = rbtree.Len()
	fmt.Printf("node count %d, dups %d \n", nodeCount, dups)
}
