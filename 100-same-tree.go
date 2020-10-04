import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	pArray := treeToArray(p)
	qArray := treeToArray(q)

	if len(pArray) != len(qArray) {
		return false
	}

	for i := 0; i < len(pArray); i++ {
		if pArray[i] != qArray[i] {
			return false
		}
	}

	return true
}

func treeToArray(t *TreeNode) []string {
	result := []string{}
	if t == nil {
		return result
	}
	result = append(result, strconv.Itoa(t.Val))
	if t.Left == nil {
		result = append(result, "null")
	} else {
		leftResults := treeToArray(t.Left)
		for _, r := range leftResults {
			result = append(result, r)
		}
	}

	if t.Right == nil {
		result = append(result, "null")
	} else {
		rightResults := treeToArray(t.Right)
		for _, r := range rightResults {
			result = append(result, r)
		}
	}

	return result
}