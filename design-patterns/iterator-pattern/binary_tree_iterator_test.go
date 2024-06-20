package iterator_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInPlaceTraversalWithForLoop(t *testing.T) {
	root := &Node{Value: 4}
	root.Left = createNode(root, 2 , 1, 3)
	right := createNode(root, 6, 5, 7)
	root.Right = right
	right = right.Right
	right.Left = createNode(right, 8, 0, 0)
	right.Right = createNode(right, 9, 0,0)

	iterator := NewBinaryTreeIterator(root, ForLoop)
	var result []int
	for val := iterator.Next(); val != nil; val = iterator.Next() {
		result = append(result, val.Value)
	}

	assert.Equal(t, []int{1,2,3,4,5,6,8,7,9}, result)
}

func TestInPlaceTraversalWithRecursion(t *testing.T) {
	root := &Node{Value: 4}
	root.Left = createNode(root, 2 , 1, 3)
	right := createNode(root, 6, 5, 7)
	root.Right = right
	right = right.Right
	right.Left = createNode(right, 8, 0, 0)
	right.Right = createNode(right, 9, 0,0)

	iterator := NewBinaryTreeIterator(root, Recursive)
	var result []int
	for val := iterator.Next(); val != nil; val = iterator.Next() {
		result = append(result, val.Value)
	}

	assert.Equal(t, []int{1,2,3,4,5,6,8,7,9}, result)
}

func createNode(parent *Node, value int, left int, right int) *Node {
	node := &Node{Value: value}
	node.Value = value

	if left != 0 {
		leftNode := &Node{Value: left, Parent: node}
		node.Left = leftNode
	}

	if right != 0 {
		rightNode := &Node{Value: right, Parent: node}
		node.Right = rightNode
	}

	node.Parent = parent;

	return node
}
