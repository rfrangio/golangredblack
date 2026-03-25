package redblack

type Comparator[K any] func(a, b K) int

type Color int
type WalkOrder int

const (
	RED Color = 1 + iota
	BLACK
)

const (
	PreOrder WalkOrder = 1 + iota
	InOrder
	PostOrder
)

type node[K any, V any] struct {
	color               Color
	left, right, parent *node[K, V]
	key                 K
	value               V
}

type Tree[K any, V any] struct {
	root       *node[K, V]
	sentinel   *node[K, V]
	cmp        Comparator[K]
	bufferPool chan *node[K, V]
	maxPool    uint32
	len        int
}

func New[K any, V any](cmp Comparator[K], maxPool uint32) *Tree[K, V] {
	tree := &Tree[K, V]{
		cmp:        cmp,
		bufferPool: make(chan *node[K, V], maxPool),
		maxPool:    maxPool,
	}

	sentinel := &node[K, V]{color: BLACK}
	sentinel.parent = sentinel
	sentinel.left = sentinel
	sentinel.right = sentinel

	tree.sentinel = sentinel
	tree.root = sentinel
	return tree
}

func (tree *Tree[K, V]) Len() int {
	return tree.len
}

func (tree *Tree[K, V]) Insert(key K, value V) bool {
	target := tree.alloc()
	target.left = tree.sentinel
	target.right = tree.sentinel
	target.parent = tree.sentinel
	target.key = key
	target.value = value

	if !tree.treeInsert(target) {
		tree.free(target)
		return false
	}

	target.color = RED
	for target != tree.root && target.parent.color == RED {
		if target.parent == target.parent.parent.left {
			uncle := target.parent.parent.right
			if uncle.color == RED {
				target.parent.color = BLACK
				uncle.color = BLACK
				target.parent.parent.color = RED
				target = target.parent.parent
			} else if target == target.parent.right {
				target = target.parent
				tree.leftRotate(target)
			} else {
				target.parent.color = BLACK
				target.parent.parent.color = RED
				tree.rightRotate(target.parent.parent)
			}
		} else {
			uncle := target.parent.parent.left
			if uncle.color == RED {
				target.parent.color = BLACK
				uncle.color = BLACK
				target.parent.parent.color = RED
				target = target.parent.parent
			} else if target == target.parent.left {
				target = target.parent
				tree.rightRotate(target)
			} else {
				target.parent.color = BLACK
				target.parent.parent.color = RED
				tree.leftRotate(target.parent.parent)
			}
		}
	}

	tree.root.color = BLACK
	tree.len++
	return true
}

func (tree *Tree[K, V]) Get(key K) (V, bool) {
	found := tree.search(key)
	if found == nil {
		var zero V
		return zero, false
	}
	return found.value, true
}

func (tree *Tree[K, V]) Delete(key K) (K, V, bool) {
	target := tree.search(key)
	if target == nil {
		var zeroKey K
		var zeroValue V
		return zeroKey, zeroValue, false
	}

	deletedKey, deletedValue := target.key, target.value
	tree.deleteNode(target)
	tree.len--
	return deletedKey, deletedValue, true
}

func (tree *Tree[K, V]) Min() (K, V, bool) {
	if tree.root == tree.sentinel {
		var zeroKey K
		var zeroValue V
		return zeroKey, zeroValue, false
	}

	minimum := tree.treeMinimum(tree.root)
	return minimum.key, minimum.value, true
}

func (tree *Tree[K, V]) Max() (K, V, bool) {
	maximum := tree.treeMaximum(tree.root)
	if maximum == nil {
		var zeroKey K
		var zeroValue V
		return zeroKey, zeroValue, false
	}

	return maximum.key, maximum.value, true
}

func (tree *Tree[K, V]) RemoveMax() (K, V, bool) {
	maximum := tree.treeMaximum(tree.root)
	if maximum == nil {
		var zeroKey K
		var zeroValue V
		return zeroKey, zeroValue, false
	}

	key, value := maximum.key, maximum.value
	tree.deleteNode(maximum)
	tree.len--
	return key, value, true
}

func (tree *Tree[K, V]) Walk(order WalkOrder, fn func(K, V) bool) {
	if fn == nil {
		return
	}
	tree.walk(tree.root, order, fn)
}

func (tree *Tree[K, V]) SubtreeDepths() (int, int) {
	var leftDepth int
	var rightDepth int
	tree.getSubtreeDepths(tree.root, 0, &leftDepth, 0, &rightDepth)
	return leftDepth, rightDepth
}

func (tree *Tree[K, V]) leftRotate(target *node[K, V]) {
	y := target.right
	target.right = y.left

	if y.left != tree.sentinel {
		y.left.parent = target
	}

	y.parent = target.parent
	if target.parent == tree.sentinel {
		tree.root = y
	} else if target == target.parent.left {
		target.parent.left = y
	} else {
		target.parent.right = y
	}

	y.left = target
	target.parent = y
}

func (tree *Tree[K, V]) rightRotate(target *node[K, V]) {
	x := target.left
	target.left = x.right

	if x.right != tree.sentinel {
		x.right.parent = target
	}

	x.parent = target.parent
	if target.parent == tree.sentinel {
		tree.root = x
	} else if target == target.parent.right {
		target.parent.right = x
	} else {
		target.parent.left = x
	}

	x.right = target
	target.parent = x
}

func (tree *Tree[K, V]) treeInsert(target *node[K, V]) bool {
	parent := tree.sentinel
	current := tree.root

	for current != tree.sentinel {
		parent = current
		comparison := tree.cmp(target.key, current.key)
		switch {
		case comparison < 0:
			current = current.left
		case comparison > 0:
			current = current.right
		default:
			return false
		}
	}

	target.parent = parent
	if parent == tree.sentinel {
		tree.root = target
	} else if tree.cmp(target.key, parent.key) < 0 {
		parent.left = target
	} else {
		parent.right = target
	}

	return true
}

func (tree *Tree[K, V]) alloc() *node[K, V] {
	if tree.maxPool == 0 {
		return new(node[K, V])
	}

	select {
	case existing := <-tree.bufferPool:
		return existing
	default:
		return new(node[K, V])
	}
}

func (tree *Tree[K, V]) free(target *node[K, V]) {
	if tree.maxPool == 0 {
		return
	}

	var zeroKey K
	var zeroValue V
	target.left = tree.sentinel
	target.right = tree.sentinel
	target.parent = tree.sentinel
	target.key = zeroKey
	target.value = zeroValue
	target.color = BLACK

	select {
	case tree.bufferPool <- target:
	default:
	}
}

func (tree *Tree[K, V]) treeMinimum(target *node[K, V]) *node[K, V] {
	for target.left != tree.sentinel {
		target = target.left
	}
	return target
}

func (tree *Tree[K, V]) treeMaximum(target *node[K, V]) *node[K, V] {
	if tree.root == tree.sentinel {
		return nil
	}

	for target.right != tree.sentinel {
		target = target.right
	}
	return target
}

func (tree *Tree[K, V]) treeSuccessor(target *node[K, V]) *node[K, V] {
	if target.right != tree.sentinel {
		return tree.treeMinimum(target.right)
	}

	parent := target.parent
	for parent != tree.sentinel && target == parent.right {
		target = parent
		parent = parent.parent
	}

	return parent
}

func (tree *Tree[K, V]) search(key K) *node[K, V] {
	current := tree.root
	for current != tree.sentinel {
		comparison := tree.cmp(key, current.key)
		switch {
		case comparison < 0:
			current = current.left
		case comparison > 0:
			current = current.right
		default:
			return current
		}
	}
	return nil
}

func (tree *Tree[K, V]) deleteNode(target *node[K, V]) {
	y := target
	if target.left != tree.sentinel && target.right != tree.sentinel {
		y = tree.treeSuccessor(target)
	}

	var x *node[K, V]
	if y.left != tree.sentinel {
		x = y.left
	} else {
		x = y.right
	}

	x.parent = y.parent
	if y.parent == tree.sentinel {
		tree.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	if y != target {
		target.key = y.key
		target.value = y.value
	}

	if y.color == BLACK {
		tree.deleteFixup(x)
	}

	tree.free(y)
}

func (tree *Tree[K, V]) deleteFixup(target *node[K, V]) {
	for target != tree.root && target.color == BLACK {
		if target == target.parent.left {
			sibling := target.parent.right
			if sibling.color == RED {
				sibling.color = BLACK
				target.parent.color = RED
				tree.leftRotate(target.parent)
				sibling = target.parent.right
			}

			if sibling.left.color == BLACK && sibling.right.color == BLACK {
				sibling.color = RED
				target = target.parent
			} else if sibling.right.color == BLACK {
				sibling.left.color = BLACK
				sibling.color = RED
				tree.rightRotate(sibling)
				sibling = target.parent.right
			} else {
				sibling.color = target.parent.color
				target.parent.color = BLACK
				sibling.right.color = BLACK
				tree.leftRotate(target.parent)
				target = tree.root
			}
		} else {
			sibling := target.parent.left
			if sibling.color == RED {
				sibling.color = BLACK
				target.parent.color = RED
				tree.rightRotate(target.parent)
				sibling = target.parent.left
			}

			if sibling.right.color == BLACK && sibling.left.color == BLACK {
				sibling.color = RED
				target = target.parent
			} else if sibling.left.color == BLACK {
				sibling.right.color = BLACK
				sibling.color = RED
				tree.leftRotate(sibling)
				sibling = target.parent.left
			} else {
				sibling.color = target.parent.color
				target.parent.color = BLACK
				sibling.left.color = BLACK
				tree.rightRotate(target.parent)
				target = tree.root
			}
		}
	}

	target.color = BLACK
}

func (tree *Tree[K, V]) walk(current *node[K, V], order WalkOrder, fn func(K, V) bool) bool {
	if current == tree.sentinel {
		return true
	}

	if order == PreOrder && !fn(current.key, current.value) {
		return false
	}
	if !tree.walk(current.left, order, fn) {
		return false
	}
	if order == InOrder && !fn(current.key, current.value) {
		return false
	}
	if !tree.walk(current.right, order, fn) {
		return false
	}
	if order == PostOrder && !fn(current.key, current.value) {
		return false
	}

	return true
}

func (tree *Tree[K, V]) getSubtreeDepths(current *node[K, V], leftLevel int, leftDepth *int, rightLevel int, rightDepth *int) {
	if current == tree.sentinel {
		return
	}

	tree.getSubtreeDepths(current.left, leftLevel+1, leftDepth, 0, rightDepth)
	tree.getSubtreeDepths(current.right, 0, leftDepth, rightLevel+1, rightDepth)

	if current.left == tree.sentinel && current.right == tree.sentinel {
		if leftLevel > *leftDepth {
			*leftDepth = leftLevel
		}
		if rightLevel > *rightDepth {
			*rightDepth = rightLevel
		}
	}
}
