import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    return getPath(root)
}

func getPath(current *TreeNode) []string {
    result := []string{}
    if current==nil {
        return result
    }
    if(current.Left!=nil) {
        for _, p := range getPath(current.Left) {
            result = append(result, fmt.Sprintf("%v->%s", current.Val, p))
        }
    }
    
    if (current.Right!=nil) {
        for _, p := range getPath(current.Right) {
            result = append(result, fmt.Sprintf("%v->%s", current.Val, p))
        }
    }
    
    if (current.Left==nil && current.Right==nil) {
        result = append(result, fmt.Sprintf("%v", current.Val))
    }
    
    return result
}