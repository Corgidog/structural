package avl

import (
	"math/rand"
	"testing"
)

func TestAvlTree(t *testing.T) {
	tree := newAvlTree()
	array := []int{5, 3, 1, 8, 9, 9, 9, 10, 11, 11, 11, 2, 4, 7, 6, 12}
	for i := 0; i < 111000; i++ {
		array = append(array, rand.Intn(10000))
	}
	for _, v := range array {
		tree.Insert(v)
	}

	output := tree.leftSort()
	t.Log(output, tree.root)
	tree.print()
}

func TestNode(t *testing.T) {
	var node *node
	node = node.insert(5)

	t.Log(node)
}
