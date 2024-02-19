package binary_search_trees

func (t *Tree) LevelOrderTraversal() [][]int {
	queue := []*TreeNode{t.Node}
	levelSize := len(queue)
	var result [][]int

	for len(queue) != 0 {
		var currNodes []int
		levelSize = len(queue)
		for i := 0; i < levelSize; i += 1 {
			curr := queue[0]
			queue = queue[1:]
			currNodes = append(currNodes, curr.Data)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		result = append(result, currNodes)
	}
	return result
}

func (t *Tree) AvgOfEachLevel() []int {
	queue := []*TreeNode{t.Node}
	var result []int

	for len(queue) != 0 {
		sum := 0
		levelSize := len(queue)
		for i := 0; i < levelSize; i += 1 {
			curr := queue[0]
			queue = queue[1:]

			sum += curr.Data
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		result = append(result, sum/levelSize)
	}
	return result
}
