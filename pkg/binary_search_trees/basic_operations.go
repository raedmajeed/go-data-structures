package binary_search_trees

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Node *TreeNode
}

func (t *Tree) Insert(value int) {
	t.Node = RecInsert(t.Node, value)
}

func RecInsert(root *TreeNode, value int) *TreeNode {
	if root == nil {
		newRoot := &TreeNode{Data: value}
		return newRoot
	}

	if value < root.Data {
		root.Left = RecInsert(root.Left, value)
	} else if value > root.Data {
		root.Right = RecInsert(root.Right, value)
	}
	return root
}

func (t *Tree) BFS() {
	queue := []*TreeNode{t.Node}

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		data := curr.Data
		fmt.Println(data, " -> ")

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
}

func TreeOperations() {
	tree := &Tree{}
	tree.Insert(5)
	tree.Insert((6))
	tree.Insert((3))
	tree.Insert((7))
	tree.Insert((4))
	tree.Insert((2))
	result1 := tree.LevelOrderTraversal()
	result2 := tree.AvgOfEachLevel()
	fmt.Println(result1, result2)
}
