import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root==nil || (root.Left==nil&&root.Right==nil) {
        return true
    }

    if ((root.Left==nil && root.Right!=nil)||(root.Left!=nil && root.Right==nil)){
        return false
    }

    leftToRight := treeToArrayLeftToRight(root.Left)
    rightToLeft := treeToArrayRightToLeft(root.Right)

    for i:=0;i<len(leftToRight);i++ {
        if leftToRight[i]!=rightToLeft[i] {
            return false
        }
    }

    return true
}

func treeToArrayLeftToRight(t *TreeNode) []string {
    resultArray := []string{}
    if t==nil {
        return resultArray
    }

    resultArray = append(resultArray, strconv.Itoa(t.Val))

    if(t.Left!=nil) {
        leftResultArray := treeToArrayLeftToRight(t.Left)
        for _, v:=range leftResultArray {
            resultArray = append(resultArray, v)
        }
    } else {
        resultArray = append(resultArray, "null")
    }

    if(t.Right!=nil) {
        rightResultArray := treeToArrayLeftToRight(t.Right)
        for _, v:=range rightResultArray {
            resultArray = append(resultArray, v)
        }
    } else {
        resultArray = append(resultArray, "null")
    }

    return resultArray
}

func treeToArrayRightToLeft(t *TreeNode) []string {
    resultArray := []string{}
    if t==nil {
        return resultArray
    }

    resultArray = append(resultArray, strconv.Itoa(t.Val))

    if(t.Right!=nil) {
        rightResultArray := treeToArrayRightToLeft(t.Right)
        for _, v:=range rightResultArray {
            resultArray = append(resultArray, v)
        }
    } else {
        resultArray = append(resultArray, "null")
    }

    if(t.Left!=nil) {
        leftResultArray := treeToArrayRightToLeft(t.Left)
        for _, v:=range leftResultArray {
            resultArray = append(resultArray, v)
        }
    } else {
        resultArray = append(resultArray, "null")
    }

    return resultArray
}