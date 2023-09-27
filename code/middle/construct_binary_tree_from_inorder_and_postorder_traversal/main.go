package construct_binary_tree_from_inorder_and_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
buildTree is a function to build a binary tree from inorder and postorder traversal.
Example: inorder = [9, 3, 15, 20, 7], postorder = [9, 15, 7, 20, 3]
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	leftSubTreeLen := findIndex(inorder, postorder[len(postorder)-1])
	root.Left = buildTree(inorder[:leftSubTreeLen], postorder[:leftSubTreeLen])
	root.Right = buildTree(inorder[leftSubTreeLen+1:], postorder[leftSubTreeLen:len(postorder)-1])

	return root
}

func findIndex(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}

	return -1
}
