package main
// Author: Robert B Frangioso

import (
         "fmt"
         "./redblack"
        )

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
    rbtree := redblack.CreateNewRedBlackTree(cmp)

	for i := 0; i < 10000000; i++ {
		rbtree.Insert(i,i)
	}

	var node_count int
	node_count = rbtree.DoNodeCount();
	fmt.Printf("node count %d \n", node_count);
	var ldepth, rdepth int
	rbtree.DoGetSubtreeDepths(&ldepth, &rdepth)
	fmt.Printf("depths %d, %d \n", ldepth, rdepth);

	for	i := 0; i < 10000000; i++{
		rbtree.RemoveMaximum()
	}

	node_count = rbtree.DoNodeCount();
	fmt.Printf("node count %d \n", node_count);
}
