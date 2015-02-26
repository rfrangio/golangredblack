package redblack 
// Author: Robert B Frangioso
import "fmt"

type Comparator func(key1 interface{}, key2 interface{}) int

type Color int
type WalkType int

const (  
	RED Color = 1 + iota  
	BLACK 
)

const (
	SILENT WalkType = 1 + iota
	PREORDER
	INORDER
	POSTORDER
)

type NODE struct {
	color Color 
	left_p, right_p, parent_p  *NODE
	key interface{}
	value interface{}
}

type RedBlackTree struct {
	m_root_p *NODE 
	m_sentinel_p *NODE
	m_sentinel NODE
	cmp_p  Comparator
}

func ConstructNODE(key interface{}, value interface{}) *NODE {
	node_p := new (NODE)
	node_p.key = key
	node_p.value = value
	return node_p
}

func ConstructRedBlackTree(cmp_p Comparator)  *RedBlackTree {
	tree_p := new (RedBlackTree)
	tree_p.m_sentinel.color = BLACK
	tree_p.m_sentinel.parent_p = nil
	tree_p.m_sentinel.left_p = nil
	tree_p.m_sentinel.right_p = nil
	tree_p.m_sentinel.key = nil
	tree_p.m_sentinel.value = nil
	tree_p.m_sentinel_p = &tree_p.m_sentinel
	tree_p.m_root_p = &tree_p.m_sentinel
	tree_p.m_root_p.parent_p = tree_p.m_sentinel_p
	tree_p.m_root_p.left_p = tree_p.m_sentinel_p
	tree_p.m_root_p.right_p = tree_p.m_sentinel_p
	tree_p.cmp_p = cmp_p
	return tree_p
}

func (tree_p *RedBlackTree) LeftRotate(target_p *NODE) {

	var y_p *NODE
	tree_p.m_sentinel_p.parent_p = target_p

	y_p = target_p.right_p

	target_p.right_p = y_p.left_p

	if y_p.left_p != tree_p.m_sentinel_p {
		y_p.left_p.parent_p = target_p
	}

	y_p.parent_p = target_p.parent_p

	if target_p.parent_p == tree_p.m_sentinel_p {
		tree_p.m_root_p = y_p

	} else if target_p == target_p.parent_p.left_p {
		target_p.parent_p.left_p = y_p
	} else {
		target_p.parent_p.right_p = y_p
	}

	y_p.left_p = target_p
	target_p.parent_p = y_p
}

func (tree_p *RedBlackTree) RightRotate(target_p *NODE) {

	tree_p.m_sentinel_p.parent_p = target_p

	var x_p *NODE
	x_p = target_p.left_p

	target_p.left_p = x_p.right_p

	if x_p.right_p != tree_p.m_sentinel_p {
		x_p.right_p.parent_p = target_p
	}

	x_p.parent_p = target_p.parent_p

	if target_p.parent_p == tree_p.m_sentinel_p {
		tree_p.m_root_p = x_p

	} else if target_p == target_p.parent_p.right_p {
		target_p.parent_p.right_p = x_p
	} else {
		target_p.parent_p.left_p = x_p
	}

	x_p.right_p = target_p
	target_p.parent_p = x_p
}

func (tree_p *RedBlackTree) TreeInsert(target_p *NODE) int {

	var y_p, trv_p	*NODE
	var ret int

	y_p = tree_p.m_sentinel_p
	trv_p = tree_p.m_root_p

	tree_p.m_sentinel_p.parent_p = tree_p.m_sentinel_p

	for trv_p != tree_p.m_sentinel_p {
		y_p = trv_p
		ret = tree_p.cmp_p(target_p.key, trv_p.key)
		if ret == -1 {
			trv_p = trv_p.left_p
		} else if ret == 1 {
			trv_p = trv_p.right_p
		} else {
			return 0
		}
	}

	target_p.parent_p = y_p
	if y_p == tree_p.m_sentinel_p {
		tree_p.m_root_p = target_p

	} else if tree_p.cmp_p(target_p.key, y_p.key) == -1 {
		y_p.left_p = target_p
	} else {
		y_p.right_p = target_p
	}
	return 1

}

func (tree_p *RedBlackTree) Insert(key interface{}, value interface{}) int {

	var target_p *NODE
	var temp_p *NODE

	target_p = new (NODE)
	temp_p = tree_p.m_sentinel_p

	target_p.left_p = tree_p.m_sentinel_p
	target_p.right_p = tree_p.m_sentinel_p
	target_p.parent_p = tree_p.m_sentinel_p

	target_p.key = key
	target_p.value = value
	tree_p.m_sentinel_p.parent_p = target_p

	if tree_p.TreeInsert(target_p) == 0 {
		target_p = nil
		return 0
	}
	tree_p.m_sentinel_p.parent_p = target_p

	target_p.color = RED

	for target_p != tree_p.m_root_p && target_p.parent_p.color == RED {

		if target_p.parent_p == target_p.parent_p.parent_p.left_p {

			temp_p = target_p.parent_p.parent_p.right_p

			if temp_p.color == RED {
				target_p.parent_p.color = BLACK
				temp_p.color = BLACK
				target_p.parent_p.parent_p.color = RED
				target_p = target_p.parent_p.parent_p
			} else if target_p == target_p.parent_p.right_p {
				target_p = target_p.parent_p
				tree_p.LeftRotate(target_p)
			} else {
				target_p.parent_p.color = BLACK
				target_p.parent_p.parent_p.color = RED
				tree_p.RightRotate(target_p.parent_p.parent_p)
			}

		} else {

			temp_p = target_p.parent_p.parent_p.left_p

			if temp_p.color == RED {
				target_p.parent_p.color = BLACK
				temp_p.color = BLACK
				target_p.parent_p.parent_p.color = RED
				target_p = target_p.parent_p.parent_p
			} else if target_p == target_p.parent_p.left_p {
				target_p = target_p.parent_p
				tree_p.RightRotate(target_p)
			} else {
				target_p.parent_p.color = BLACK
				target_p.parent_p.parent_p.color = RED
				tree_p.LeftRotate(target_p.parent_p.parent_p)
			}

		}
	}

	tree_p.m_root_p.color = BLACK
	return 1
}

func (tree_p *RedBlackTree) TreeMinimum(target_p *NODE) *NODE {

	for target_p.left_p != tree_p.m_sentinel_p {
		target_p = target_p.left_p
	}

	return target_p
}

func (tree_p *RedBlackTree) Minimum() (interface{}, interface{}) {

	var target_p *NODE
	var ret_key, ret_val interface{}	

	target_p = tree_p.m_root_p
	

	for target_p.left_p != tree_p.m_sentinel_p {
		target_p = target_p.left_p
	}

	ret_key = target_p.key
	ret_val = target_p.value

	return ret_key, ret_val 
}

func (tree_p *RedBlackTree) TreeMaximum(target_p *NODE) *NODE {

	if tree_p.m_root_p == tree_p.m_sentinel_p {
		return nil
	}

	for target_p.right_p != tree_p.m_sentinel_p {
		target_p = target_p.right_p
	}

	return target_p
}

func (tree_p *RedBlackTree) RemoveMaximum() (interface{}, interface{}) {

	var node_p *NODE
	var	ret_key interface{}
	var	ret_value interface{}

	node_p = tree_p.TreeMaximum(tree_p.m_root_p)

	if node_p == nil {
		return nil, nil
	}

	ret_key = node_p.key
	ret_value = node_p.value

	tree_p.TreeDelete(node_p)

	return ret_key, ret_value
}

func (tree_p *RedBlackTree) Maximum() (interface{}) {

	var target_p *NODE
	target_p = tree_p.m_root_p

	for target_p.right_p != tree_p.m_sentinel_p {
		target_p = target_p.right_p
	}

	return target_p.key
}

func (tree_p *RedBlackTree) TreeSuccessor(target_p *NODE) *NODE {

	var trv_p *NODE

	if target_p != tree_p.m_sentinel_p {
		return tree_p.TreeMinimum(target_p.right_p)
	}

	trv_p = target_p.parent_p
	for trv_p != tree_p.m_sentinel_p && target_p == trv_p.right_p {
		target_p = trv_p
		trv_p = trv_p.parent_p
	}

	return trv_p
}

func (tree_p *RedBlackTree) Search(key interface{}) *NODE {

	var trv_p	*NODE
	trv_p = tree_p.m_root_p

	for trv_p != tree_p.m_sentinel_p && tree_p.cmp_p(key, trv_p.key) != 0 {
		if tree_p.cmp_p(key, trv_p.key) == -1 {
			trv_p = trv_p.left_p
		} else {
			trv_p = trv_p.right_p
		}
	}

	if trv_p == tree_p.m_sentinel_p {
		return nil
	} else {
		return trv_p
	}
}

func (tree_p *RedBlackTree) Delete(key interface{}) (interface{}, interface{}) {

	var target_p, y_p, x_p *NODE
	var ret_key interface{}
	var ret_value interface{}

	target_p = tree_p.Search(key)

	if (target_p == nil) {
		return nil, nil
	}

	tree_p.m_sentinel_p.parent_p = target_p

	if target_p.left_p == tree_p.m_sentinel_p || target_p.right_p == tree_p.m_sentinel_p {
		y_p = target_p
	} else {
		y_p = tree_p.TreeSuccessor(target_p)
	}

	if y_p.left_p != tree_p.m_sentinel_p {
		x_p = y_p.left_p
	} else {
		x_p = y_p.right_p
	}

	x_p.parent_p = y_p.parent_p

	if y_p.parent_p == tree_p.m_sentinel_p {
		tree_p.m_root_p = x_p
	} else if y_p == y_p.parent_p.left_p {
		y_p.parent_p.left_p = x_p
	} else {
		y_p.parent_p.right_p = x_p
	}

	if y_p != target_p {
		ret_key = target_p.key
		ret_value = target_p.value
	} else {
		ret_key = y_p.key
		ret_value = y_p.value
	}

	if (y_p.color == BLACK) {
		tree_p.DeleteFixup(x_p)
	}
	y_p = nil

	return ret_key, ret_value

}

func (tree_p *RedBlackTree) TreeDelete(target_p *NODE) (interface{}, interface{}) {

	var y_p, x_p *NODE
	var	ret_key interface{}
	var	ret_value interface{}

	tree_p.m_sentinel_p.parent_p = target_p

	if target_p == tree_p.m_sentinel_p {
		return nil, nil
	}

	if target_p.left_p == tree_p.m_sentinel_p || target_p.right_p == tree_p.m_sentinel_p {
		y_p = target_p
	} else {
		y_p = tree_p.TreeSuccessor(target_p)
	}

	if y_p.left_p != tree_p.m_sentinel_p {
		x_p = y_p.left_p
	} else {
		x_p = y_p.right_p
	}

	x_p.parent_p = y_p.parent_p

	if y_p.parent_p == tree_p.m_sentinel_p {
		tree_p.m_root_p = x_p
	} else if y_p == y_p.parent_p.left_p {
		y_p.parent_p.left_p = x_p
	} else {
		y_p.parent_p.right_p = x_p
	}

	if (y_p != target_p) {
		ret_key = target_p.key
		ret_value = target_p.value

		target_p.key = y_p.key
		target_p.value = y_p.value
	} else {
		ret_key = y_p.key
		ret_value = y_p.value
	}

	if (y_p.color == BLACK) {
		tree_p.DeleteFixup(x_p)
	}
	y_p = nil

	return ret_key, ret_value
}

func (tree_p *RedBlackTree) DeleteFixup(target_p *NODE) {

	var w_p	*NODE
	tree_p.m_sentinel_p.parent_p = target_p

	for target_p != tree_p.m_root_p && target_p.color == BLACK {


		if target_p == target_p.parent_p.left_p {
			w_p = target_p.parent_p.right_p
			tree_p.m_sentinel_p.parent_p = w_p

			if w_p.color == RED {
				w_p.color = BLACK
				target_p.parent_p.color = RED

				tree_p.LeftRotate(target_p.parent_p)

				w_p = target_p.parent_p.right_p
				tree_p.m_sentinel_p.parent_p = w_p
			}

			if w_p.left_p.color == BLACK && w_p.right_p.color == BLACK {
				w_p.color = RED
				target_p = target_p.parent_p
			} else if w_p.right_p.color == BLACK {
				w_p.left_p.color = BLACK
				w_p.color = RED

				tree_p.RightRotate(w_p)

				w_p = target_p.parent_p.right_p
				tree_p.m_sentinel_p.parent_p = w_p
			} else {
				w_p.color = target_p.parent_p.color
				target_p.parent_p.color = BLACK
				w_p.right_p.color = BLACK

				tree_p.LeftRotate(target_p.parent_p)

				target_p = tree_p.m_root_p
				tree_p.m_sentinel_p.parent_p = target_p
			}

		} else {

			w_p = target_p.parent_p.left_p
			tree_p.m_sentinel_p.parent_p = target_p

			if w_p.color == RED {
				w_p.color = BLACK
				target_p.parent_p.color = RED

				tree_p.RightRotate(target_p.parent_p)

				w_p = target_p.parent_p.left_p
				tree_p.m_sentinel_p.parent_p = w_p

			}

			if w_p.right_p.color == BLACK && w_p.left_p.color == BLACK {
				w_p.color = RED
				target_p = target_p.parent_p
			} else if w_p.left_p.color == BLACK {
				w_p.right_p.color = BLACK
				w_p.color = RED

				tree_p.LeftRotate(w_p)

				w_p = target_p.parent_p.left_p
				tree_p.m_sentinel_p.parent_p = w_p
			} else {

				w_p.color = target_p.parent_p.color
				target_p.parent_p.color = BLACK
				w_p.left_p.color = BLACK

				tree_p.RightRotate(target_p.parent_p)

				target_p = tree_p.m_root_p
				tree_p.m_sentinel_p.parent_p = target_p
			}

		}

	}

	target_p.color = BLACK

}

func (tree_p *RedBlackTree) NodeCount(node_p *NODE, level int) int {

	var	i int
	i = 1	

	if (node_p == tree_p.m_sentinel_p) {
		return 0 
	}

	i += tree_p.NodeCount(node_p.left_p, level + 1)
	i += tree_p.NodeCount(node_p.right_p, level + 1)

	return i
}

func (tree_p *RedBlackTree) DoNodeCount() int {

	var ret int 
	ret = 0

	ret = tree_p.NodeCount(tree_p.m_root_p, 0)

	return ret

}

func (tree_p *RedBlackTree) Walk(node_p *NODE, level int, walk_type WalkType) {

	if (node_p == tree_p.m_sentinel_p) {
		return
	}
	
	if walk_type == PREORDER {
		fmt.Printf("key %v, level %d\n", node_p.key, level) 
	}

	tree_p.Walk(node_p.left_p, level + 1, walk_type)

	if walk_type == INORDER {
		fmt.Printf("key %v, level %d\n", node_p.key, level) 
	}

	tree_p.Walk(node_p.right_p, level + 1, walk_type)

	if walk_type == POSTORDER {
		fmt.Printf("key %v, level %d\n", node_p.key, level) 
	}
}

func (tree_p *RedBlackTree) DoWalk(walk_type WalkType) {

	tree_p.Walk(tree_p.m_root_p, 0, walk_type)

}

func (tree_p *RedBlackTree) DoGetSubtreeDepths(left_tree *int, right_tree *int) {

	tree_p.GetSubtreeDepths(tree_p.m_root_p, 0, left_tree, 0, right_tree)
}

func (tree_p *RedBlackTree) GetSubtreeDepths(node_p *NODE, left_level int, left_tree *int, right_level int,right_tree *int) {


	if (node_p == tree_p.m_sentinel_p) {
		return
	}

	tree_p.GetSubtreeDepths(node_p.left_p, left_level + 1, left_tree, 0, right_tree)
	tree_p.GetSubtreeDepths(node_p.right_p, 0, left_tree, right_level + 1, right_tree)

	if node_p.left_p ==tree_p.m_sentinel_p && node_p.right_p ==tree_p.m_sentinel_p {
		if left_level > *left_tree {
			*left_tree = left_level
		}
		if right_level > *right_tree {
			*right_tree = right_level
		}
	}
}

