package leetcode

func traversal_inorder(root *TreeNode, rst []int) []int {
	if root == nil {
		return rst
	}

	rst = traversal_inorder(root.Left, rst)
	rst = append(rst, root.Val)
	rst = traversal_inorder(root.Right, rst)

	return rst
}

func traversal_preorder(root *TreeNode, rst []int) []int {
	if root == nil {
		return rst
	}

	rst = append(rst, root.Val)
	rst = traversal_preorder(root.Left, rst)
	rst = traversal_preorder(root.Right, rst)

	return rst
}

func traversal_postorder(root *TreeNode, rst []int) []int {
	if root == nil {
		return rst
	}

	rst = traversal_preorder(root.Left, rst)
	rst = traversal_preorder(root.Right, rst)
	rst = append(rst, root.Val)

	return rst
}

func traversal_recursion(root *TreeNode, rst []int) []int {
	if root == nil {
		return rst
	}

	curr, stack := root, []*TreeNode{}
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curr = curr.Right
	}
	return rst
}
