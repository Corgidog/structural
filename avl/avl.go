package avl

import (
	"bytes"
	"fmt"
)

type avlTree struct {
	root *node
}

func newAvlTree() *avlTree {
	return &avlTree{}
}

func (tree *avlTree) Insert(data int) {
	tree.root = tree.root.insert(data)
}

func (tree *avlTree) leftSort() []int {
	outputs := make([]int, 0, 10)
	outputs = tree.root.leftSort(outputs)
	return outputs
}

func (tree *avlTree) print() {
	tree._print([]*node{tree.root})
}

func (tree *avlTree) _print(nodes []*node) {
	children := make([]*node, 0, 2*len(nodes))
	var buf bytes.Buffer
	index := 0
	for _, n := range nodes {
		if n != nil {
			index++
			buf.WriteString(fmt.Sprintf("%d:h:%d,v:%d  ", index, n.getHeight(), n.data))
			children = append(children, n.left)
			children = append(children, n.right)
		}
	}
	fmt.Println(buf.String())
	if len(children) > 0 {
		tree._print(children)
	}
}

type node struct {
	left   *node
	right  *node
	height int
	data   int
}

func (n *node) insert(data int) *node {
	if n == nil {
		return &node{data: data}
	}

	if data < n.data {
		n.left = n.left.insert(data)
		n.setHeight()
		if n.left.getHeight()-n.right.getHeight() == 2 {
			if data < n.left.data {
				n = n.rightRotation()
			} else if data > n.left.data {
				n = n.leftRightRotation()
			}
		}
	} else if data > n.data {
		n.right = n.right.insert(data)
		n.setHeight()
		if n.right.getHeight()-n.left.getHeight() == 2 {
			if data > n.right.data {
				n = n.leftRotation()
			} else if data < n.right.data {
				n = n.rightLeftRotation()
			}
		}
	}

	return n
}

func (n *node) setHeight() {
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
}

// 左旋
func (n *node) leftRotation() *node {
	right := n.right
	n.right = right.left
	right.left = n
	n.setHeight()
	right.setHeight()
	return right
}

func (n *node) rightRotation() *node {
	left := n.left
	n.left = left.right
	left.right = n
	left.setHeight()
	n.setHeight()
	return left
}

func (n *node) rightLeftRotation() *node {
	n.right = n.right.rightRotation()
	return n.leftRotation()
}

func (n *node) leftRightRotation() *node {
	n.left = n.left.leftRotation()
	return n.rightRotation()
}

func (n *node) getHeight() int {
	if n == nil {
		return -1
	}

	return n.height
}

func (n *node) leftSort(output []int) []int {
	if n != nil {
		output = n.left.leftSort(output)
		output = append(output, n.data)
		output = n.right.leftSort(output)
	}

	return output
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
