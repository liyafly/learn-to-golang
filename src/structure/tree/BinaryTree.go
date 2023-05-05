package tree

func BuildTreeIter(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 使用栈来构建树
	stack := make([]*TreeNode, 0)
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	stack = append(stack, root)
	preorderIndex := 1
	inorderIndex := 0
	var prev *TreeNode
	for preorderIndex < len(preorder) {
		// 遍历栈中的节点
		for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
			prev = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			inorderIndex++
		}
		// 创建当前节点，并将其添加到父节点的左侧或者右侧
		node := &TreeNode{Val: preorder[preorderIndex]}
		if prev != nil {
			prev.Right = node
		} else {
			stack[len(stack)-1].Left = node

		}
		stack = append(stack, node)
		preorderIndex++
		prev = nil
	}
	return root
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 前序遍历的第一个节点是根节点
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	var i int
	for i = range inorder {
		if inorder[i] == rootVal {
			break
		}
	}
	// 递归构建左子树和右子树
	root.Left = BuildTree(preorder[1:i+1], inorder[:i])
	root.Right = BuildTree(preorder[i+1:], inorder[i+1:])
	return root

}
