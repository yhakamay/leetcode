package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//root1 := TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}}
	////root2 := TreeNode{2, 1, 3, null, 4, null, 7}
	//root2 := TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3, Left: nil}, 4, null, 7}
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	newRoot := TreeNode{Val: root1.Val + root2.Val}
	newRoot.Left = mergeTrees(root1.Left, root2.Left)
	newRoot.Right = mergeTrees(root1.Right, root2.Right)

	return &newRoot
}
