package hw10tree

import (
	"fmt"
)

type Node struct {
	key    int
	parent *Node
	left   *Node
	right  *Node
}

// SimpleTree простое несбалансированное двоичное дерево.
type SimpleTree struct {
	root *Node
}

// Insert вставить узел с ключом x. Если такой ключ существует то ничего не делаем.
func (t *SimpleTree) Insert(x int) {
	current, parent := t.search(x)

	if current != nil {
		return
	}

	node := &Node{
		key:    x,
		parent: parent,
	}

	if parent == nil {
		t.root = node
		return
	}

	if node.key < parent.key {
		parent.left = node
		return
	}

	parent.right = node
}

// Search найден или нет узел с ключом равным х.
func (t *SimpleTree) Search(x int) bool {
	node, _ := t.search(x)
	return node != nil
}

func (t *SimpleTree) removeLeaf(node *Node) {
	if node == t.root {
		t.root = nil
		return
	}

	if t.isLeftChild(node) {
		node.parent.left = nil
		return
	}

	node.parent.right = nil
}

func (t *SimpleTree) removeHasOnlyRightChild(node *Node) {
	if t.root == node {
		t.root = node.right
		node.right.parent = nil
		return
	}

	parent := node.parent
	if t.isLeftChild(node) {
		parent.left = node.right
	} else {
		parent.right = node.right
	}
	node.right.parent = parent
}

func (t *SimpleTree) removeHasOnlyLeftChild(node *Node) {
	if t.root == node {
		t.root = node.left
		node.left.parent = nil
		return
	}

	parent := node.parent
	if t.isLeftChild(node) {
		parent.left = node.left
	} else {
		parent.right = node.left
	}
	node.left.parent = parent
}

func (t *SimpleTree) removeHasBothChild(node *Node) {
	maxLeftNode := t.getMaxLeft(node)

	if t.isLeaf(maxLeftNode) {
		t.removeLeaf(maxLeftNode)
	} else {
		t.removeHasOnlyLeftChild(maxLeftNode)
	}

	maxLeftNode.left, maxLeftNode.right, maxLeftNode.parent = node.left, node.right, node.parent
	if maxLeftNode.left != nil {
		maxLeftNode.left.parent = maxLeftNode
	}
	if maxLeftNode.right != nil {
		maxLeftNode.right.parent = maxLeftNode
	}

	if maxLeftNode.parent != nil {
		if t.isLeftChild(node) {
			maxLeftNode.parent.left = maxLeftNode
		} else {
			maxLeftNode.parent.right = maxLeftNode
		}
	}

	if node == t.root {
		t.root = maxLeftNode
	}
}

// Remove удалить узел с ключом x.
func (t *SimpleTree) Remove(x int) {
	node, _ := t.search(x)
	if node == nil {
		return
	}

	if t.isLeaf(node) {
		t.removeLeaf(node)
		return
	}

	if t.hasOnlyRightChild(node) {
		t.removeHasOnlyRightChild(node)
		return
	}

	if t.hasOnlyLeftChild(node) {
		t.removeHasOnlyLeftChild(node)
		return
	}

	t.removeHasBothChild(node)

	node.parent, node.left, node.right = nil, nil, nil
}

// Sort Вывести в отстортированном виде ключи дерева по возрастанию.
func (t *SimpleTree) Sort() {
	t.print(t.root)
}

func (t *SimpleTree) print(node *Node) {
	if node == nil {
		return
	}
	t.print(node.left)
	fmt.Printf("%d ", node.key)
	t.print(node.right)
}

// search Возвращает узел ключ которого равен x и его предка.
func (t *SimpleTree) search(x int) (*Node, *Node) {
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

// getMaxLeft вернуть узел с максимальным ключом из левого поддерева.
func (t *SimpleTree) getMaxLeft(node *Node) *Node {
	current := node.left
	for {
		if current.right == nil {
			return current
		}
		current = current.right
	}
}

// isLeaf является ли узел листом дерева.
func (t *SimpleTree) isLeaf(node *Node) bool {
	return node.left == nil && node.right == nil
}

// isLeftChild является ли узел левым ребенком.
func (t *SimpleTree) isLeftChild(node *Node) bool {
	return node.parent.left == node
}

// hasOnlyRightChild имеет ли узел только правого ребенка.
func (t *SimpleTree) hasOnlyRightChild(node *Node) bool {
	return node.left == nil && node.right != nil
}

// hasOnlyLeftChild имеет ли узел только левого ребенка.
func (t *SimpleTree) hasOnlyLeftChild(node *Node) bool {
	return node.left != nil && node.right == nil
}
