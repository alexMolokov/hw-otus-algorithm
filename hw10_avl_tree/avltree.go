package hw10tree

import "fmt"

type AvlNode struct {
	key    int
	height int
	parent *AvlNode
	left   *AvlNode
	right  *AvlNode
}

type AvlTree struct {
	root *AvlNode
}

func (t *AvlTree) search(x int) (*AvlNode, *AvlNode) {
	if t.root == nil {
		return nil, nil
	}

	current := t.root
	for {
		if x == current.key {
			return current, current.parent
		}

		if x < current.key {
			if current.left == nil {
				return nil, current
			}
			current = current.left
			continue
		}

		if current.right == nil {
			return nil, current
		}
		current = current.right
	}
}

func (t *AvlTree) Remove(x int) {
	node, _ := t.search(x)
	if node == nil {
		return
	}

	if t.isLeaf(node) {
		t.removeLeaf(node)
		return
	}

	if t.hasOneChild(node) {
		t.removeHasOnlyOneChild(node)
		return
	}

	t.removeHasBothChild(node)
}

// getMinRight вернуть узел с минимальным ключом из правого поддерева.
func (t *AvlTree) getMinRight(node *AvlNode) *AvlNode {
	current := node.right
	for {
		if current.left == nil {
			return current
		}
		current = current.left
	}
}

func (t *AvlTree) removeHasBothChild(node *AvlNode) {
	next := t.getMinRight(node)
	next.key, node.key = node.key, next.key

	if t.isLeaf(next) {
		t.removeLeaf(next)
		return
	}

	t.removeHasOnlyOneChild(next)
}

func (t *AvlTree) removeLeaf(node *AvlNode) {
	if node == t.root {
		t.root = nil
		return
	}

	if t.isLeftChild(node) {
		node.parent.left = nil
	} else {
		node.parent.right = nil
	}

	t.fixHeight(node.parent)
	node.left, node.right, node.parent = nil, nil, nil
}

func (t *AvlTree) removeHasOnlyOneChild(node *AvlNode) {
	child := node.left
	if child == nil {
		child = node.right
	}

	parent := node.parent
	if parent == nil {
		t.root = child
		return
	}

	if t.isLeftChild(node) {
		parent.left = child
	} else {
		parent.right = child
	}
	child.parent = parent

	t.fixHeight(parent)
	node.left, node.right, node.parent = nil, nil, nil
}

// isLeaf является ли узел листом дерева.
func (t *AvlTree) isLeaf(node *AvlNode) bool {
	return node.left == nil && node.right == nil
}

// hasOneChild ...
func (t *AvlTree) hasOneChild(node *AvlNode) bool {
	return (node.left == nil && node.right != nil) ||
		(node.right == nil && node.left != nil)
}

func (t *AvlTree) Insert(x int) {
	current, parent := t.search(x)

	if current != nil {
		return
	}

	node := &AvlNode{
		key:    x,
		parent: parent,
		height: 1,
	}

	if parent == nil {
		t.root = node
		return
	}

	if node.key < parent.key {
		parent.left = node
	} else {
		parent.right = node
	}

	t.fixHeight(parent)
}

func (t *AvlTree) fixHeight(node *AvlNode) {
	current := node
	for {
		leftHeight := getHeight(current.left)
		rightHeight := getHeight(current.right)

		maxHeight := recalculateHeight(current.left, current.right)

		current.height = maxHeight
		// если разница между высотами поддеревьев > 1 то требуется перебалансировка
		if abs(leftHeight, rightHeight) > 1 {
			current = t.reBalance(current)
		}

		if current == t.root {
			break
		}
		current = current.parent
	}
}

func (t *AvlTree) smallRightRotation(node *AvlNode) *AvlNode {
	p := node.parent
	dp := p.parent // дед
	c := node.right

	if dp != nil {
		if t.isLeftChild(p) {
			dp.left = node
		} else {
			dp.right = node
		}
	}
	node.parent = dp
	p.parent = node
	node.right = p
	p.left = c

	if c != nil {
		c.parent = p
	}

	p.height = recalculateHeight(p.left, p.right)
	node.height = recalculateHeight(node.left, node.right)

	if dp == nil {
		t.root = node
	}

	return node
}

// smallLeftRotation передавать элемент который будем крутить.
func (t *AvlTree) smallLeftRotation(node *AvlNode) *AvlNode {
	p := node.parent
	dp := p.parent
	c := node.left

	if dp != nil {
		if t.isLeftChild(p) {
			dp.left = node
		} else {
			dp.right = node
		}
	}

	node.parent = dp
	p.parent = node
	node.left = p
	p.right = c

	if c != nil {
		c.parent = p
	}

	p.height = recalculateHeight(p.left, p.right)
	node.height = recalculateHeight(node.right, node.left)

	if dp == nil {
		t.root = node
	}

	return node
}

func (t *AvlTree) bigRightRotation(node *AvlNode) *AvlNode {
	next := t.smallLeftRotation(node.left.right)
	return t.smallRightRotation(next)
}

func (t *AvlTree) bigLeftRotation(node *AvlNode) *AvlNode {
	next := t.smallRightRotation(node.right.left)
	return t.smallLeftRotation(next)
}

func (t *AvlTree) reBalance(node *AvlNode) *AvlNode {
	leftHeight := getHeight(node.left)
	rightHeight := getHeight(node.right)
	if abs(leftHeight, rightHeight) <= 1 {
		return node
	}

	if leftHeight > rightHeight {
		l := getHeight(node.left.left)
		c := getHeight(node.left.right)

		if l >= c {
			return t.smallRightRotation(node.left)
		}
		return t.bigRightRotation(node)
	}

	r := getHeight(node.right.right)
	c := getHeight(node.right.left)

	if r >= c {
		return t.smallLeftRotation(node.right)
	}
	return t.bigLeftRotation(node)
}

// Sort Вывести в отстортированном виде ключи дерева по возрастанию.
func (t *AvlTree) Sort() {
	t.print(t.root)
}

func (t *AvlTree) print(node *AvlNode) {
	if node == nil {
		return
	}
	t.print(node.left)
	fmt.Printf("%d ", node.key)
	t.print(node.right)
}

func (t *AvlTree) isLeftChild(node *AvlNode) bool {
	return node.parent.left == node
}

func recalculateHeight(x, y *AvlNode) int {
	xHeight := 0
	if x != nil {
		xHeight = x.height
	}

	yHeight := 0
	if y != nil {
		yHeight = y.height
	}

	if xHeight > yHeight {
		return xHeight + 1
	}
	return yHeight + 1
}

func getHeight(node *AvlNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
