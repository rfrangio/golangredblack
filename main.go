package main

// Author: Robert B Frangioso

import (
	"./redblack"
	"fmt"
	"math/rand"
	"time"
)

func cmp_c(key1 interface{}, key2 interface{}) int {
	char_key1 := key1.(string)
	char_key2 := key2.(string)

	if char_key1 == char_key2 {
		return 0
	}

	if char_key1 > char_key2 {
		return 1
	} else {
		return -1
	}
}

func cmp(key1 interface{}, key2 interface{}) int {
	int_key1 := key1.(int)
	int_key2 := key2.(int)

	if int_key1 == int_key2 {
		return 0
	}

	if int_key1 > int_key2 {
		return 1
	} else {
		return -1
	}
}

func main() {

	rbtree := redblack.ConstructRedBlackTree(cmp, 0)
	rbtreec := redblack.ConstructRedBlackTree(cmp_c, 1000)
	var num_objects int = 1939347
	var dups int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rbtreec.Insert("a", "a")

	fmt.Printf("Inserting %d sequential objects \n", num_objects)

	for i := 0; i < num_objects; i++ {
		rbtree.Insert(i, i)
	}

	var node_count int
	node_count = rbtree.DoNodeCount()
	fmt.Printf("node count %d \n", node_count)
	var ldepth, rdepth int
	rbtree.DoGetSubtreeDepths(&ldepth, &rdepth)
	fmt.Printf("depths %d, %d \n", ldepth, rdepth)

	for i := 0; i < num_objects; i++ {
		rbtree.RemoveMaximum()
	}

	node_count = rbtree.DoNodeCount()
	fmt.Printf("node count %d \n", node_count)

	fmt.Printf("Inserting %d random objects \n", num_objects)
	node_count, ldepth, rdepth = 0, 0, 0

	var entry int
	for i := 0; i < num_objects; i++ {
		entry = r.Int()
		if rbtree.Insert(entry, entry) == 0 {
			dups++
		}
	}

	node_count = rbtree.DoNodeCount()
	fmt.Printf("node count %d \n", node_count+dups)
	rbtree.DoGetSubtreeDepths(&ldepth, &rdepth)
	fmt.Printf("depths %d, %d \n", ldepth, rdepth)

	for i := 0; i < num_objects; i++ {
		rbtree.RemoveMaximum()
	}

	node_count = rbtree.DoNodeCount()
	fmt.Printf("node count %d, dups %d \n", node_count, dups)
}
