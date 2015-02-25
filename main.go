package main
// Author: Robert B Frangioso

import (
         "fmt"
         "./redblack"
        )

func cmp(key1 interface{}, key2 interface{}) int {
  intKey1 := key1.(int)
  intKey2 := key2.(int)

  if intKey1 == intKey2 {
    return 0
  } else if intKey1 > intKey2 {
    return 1
  }

  return -1
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
