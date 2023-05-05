package tree

// BuildTreeIter
/*
1. 将先序遍历数组的第一个元素作为根节点，创建一个树节点。
2. 将这个根节点入栈。
3. 遍历先序遍历数组（除了根节点），取出一个元素，创建一个树节点。
4. 判断当前节点和栈顶元素是否相等，如果相等，则说明当前节点应该是栈顶元素的左儿子；如果不相等，则说明当前节点应该是栈顶元素的右儿子。
5. 如果当前节点是栈顶元素的左儿子，则将当前节点入栈。
6. 如果当前节点是栈顶元素的右儿子，则弹出栈顶元素，直到栈顶元素的右儿子为空或者与当前节点不相等为止，再将当前节点入栈。
7. 遍历完先序遍历数组，将栈中的树节点从栈底到栈顶依次取出，分别构建左子树和右子树。
8. 返回根节点。
*/
func BuildTreeIter(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 使用栈来构建树
	stack := make([]*TreeNode, 0)
	rootVal := preorder[0]
	// 先序遍历的第一个节点是根节点
	root := &TreeNode{Val: rootVal}
	// 根节点入栈
	stack = append(stack, root)
	preorderIndex := 1
	inorderIndex := 0
	var prev *TreeNode
	// 遍历先序遍历数组，除了根节点之外的节点都入栈
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

// PreorderTraversal 前序遍历 递归
func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	left := PreorderTraversal(root.Left)
	right := PreorderTraversal(root.Right)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// PreorderTraversalIter 前序遍历 迭代
func PreorderTraversalIter(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			res = append(res, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		}
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr = pop.Right
	}
	return res
}

// InorderTraversalIter 中序遍历 迭代
func InorderTraversalIter(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, pop.Val)
		curr = pop.Right
	}
	return res
}

// InorderTraversal 中序遍历 递归
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := InorderTraversal(root.Left)
	result := []int{root.Val}
	right := InorderTraversal(root.Right)
	result = append(left, result...)
	result = append(result, right...)
	return result
}

// PostorderTraversal 后序遍历 递归
func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := PostorderTraversal(root.Left)
	right := PostorderTraversal(root.Right)
	result := []int{}
	result = append(left, right...)
	result = append(result, root.Val)
	return result
}

// PostorderTraversalIter 后序遍历 迭代
func PostorderTraversalIter(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	var res []int
	stack = append(stack, root)
	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{pop.Val}, res...)
		if pop.Left != nil {
			stack = append(stack, pop.Left)
		}
		if pop.Right != nil {
			stack = append(stack, pop.Right)
		}
	}
	return res
}

// 层次遍历
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	queue := []*TreeNode{root} // 将根节点加入队列
	// 迭代层次遍历二叉树
	for len(queue) != 0 {
		levelSize := len(queue) // 当前层的节点个数
		levelValues := make([]int, levelSize)
		// 遍历当前层的所有节点，并将它们放到 levelValues 数组中
		for i := 0; i < levelSize; i++ {
			node := queue[0]  // 取出队列头节点
			queue = queue[1:] // 弹出队列头节点
			levelValues[i] = node.Val
			// 将当前节点的子节点放入队列中
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 将当前层的节点值存到结果中
		res = append(res, levelValues)
	}
	return res
}
