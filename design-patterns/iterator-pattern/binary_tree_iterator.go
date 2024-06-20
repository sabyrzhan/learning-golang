package iterator_pattern

// Iteration over binary tree with in-place traversal
// In-place traversal starts from leftmost item, routes through parent and down to rightmost value
// For example:
//       8
//      /  \
//     7    9
// The result will be: 7-8-9
// There are 2 implementations here: with for loop and using recursion.
// In for loop:
//		the node must have reference also to parent node, alongside with left and right values. This is needed to stop
//		the loop in order not to visit the same nodes over and over again. Also as a root you must pass the left most item
//		as a starting point
// In recursion:
//		we can traverse left and right sides separately since no need of knowledge about the parents. The same function
//		is called for left side and right sides separately. You must pass the root node as a starting point as opposed
//		to the loop version

type IteratorType int
const (
	ForLoop IteratorType = iota
	Recursive
)

type Node struct {
	Value int
	Parent *Node
	Left  *Node
	Right *Node
}

type BinaryTreeIterator interface {
	Next() *Node
}

func NewBinaryTreeIterator(root *Node, itType IteratorType) BinaryTreeIterator {
	var result BinaryTreeIterator
	switch itType {
	case ForLoop:
		leftNode := root;
		for leftNode.Left != nil {
			leftNode = leftNode.Left
		}
		result = &BinaryTreeIteratorForLoop{Current: leftNode, isFirstVisit: false}
	case Recursive:
		result = &BinaryTreeIteratorWithRecursion{Current: root}
	}

	return result
}

type BinaryTreeIteratorForLoop struct {
	Current *Node
	isFirstVisit bool
}

type BinaryTreeIteratorWithRecursion struct {
	Current *Node
	result []*Node
	currentIndex int
	recursionFinished bool
}

func (it *BinaryTreeIteratorForLoop) Next() *Node {
	if it.Current == nil {
		return nil
	}

	if !it.isFirstVisit {
		it.isFirstVisit = true
		return it.Current
	}

	if it.Current.Right != nil {
		it.Current = it.Current.Right
		for it.Current.Left != nil {
			it.Current = it.Current.Left
		}

		return it.Current
	} else {
		parent := it.Current.Parent
		for parent != nil && parent.Right == it.Current {
			it.Current = parent
			parent = parent.Parent
		}
		it.Current = parent

		return it.Current
	}
}

func (it *BinaryTreeIteratorWithRecursion) nextRecursive(node *Node) {
	if node == nil {
		return
	}

	it.nextRecursive(node.Left)
	it.result = append(it.result, node)
	it.nextRecursive(node.Right)
}
func (it *BinaryTreeIteratorWithRecursion) Next() *Node {
	if !it.recursionFinished {
		it.nextRecursive(it.Current)
		it.recursionFinished = true
	}

	if len(it.result) <= it.currentIndex {
		return nil
	} else {
		node := it.result[it.currentIndex]
		it.currentIndex++
		return node
	}
}
