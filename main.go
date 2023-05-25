package main

import (
	"fmt"
	"learn-to-golang/DataStructuresAndAlgorithms/sort"
	. "learn-to-golang/DataStructuresAndAlgorithms/tree"
	"learn-to-golang/design/singleton"
	sort2 "sort"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibIter(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	array := []int{-74, 48, -20, 2, 10, -84, -5, -9, 11, -24, -91, 2, -71, 64, 63, 80, 28, -30, -58, -11, -44, -87, -22, 54, -74, -10, -55, -28, -46, 29, 10, 50, -72, 34, 26, 25, 8, 51, 13, 30, 35, -8, 50, 65, -6, 16, -2, 21, -78, 35, -13, 14, 23, -3, 26, -90, 86, 25, -56, 91, -13, 92, -25, 37, 57, -20, -69, 98, 95, 45, 47, 29, 86, -28, 73, -44, -46, 65, -84, -96, -24, -12, 72, -68, 93, 57, 92, 52, -45, -2, 85, -63, 56, 55, 12, -85, 77, -39}
	fmt.Println(array)
	sort2.Ints(array)
	sort.Quicksort(array)

	fmt.Println(array)
	aa := []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1}
	fmt.Println(aa)
	sort.Quicksort(aa)

	fmt.Println(aa)
	// 创建一棵二叉树

	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}
	root.Left = node2
	root.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7

	// 递归前序遍历
	fmt.Printf("递归前序遍历：	 =%v\n", PreorderTraversal(root))
	// 非递归前序遍历
	fmt.Printf("非递归前序遍历	 =%v\n", PreorderTraversalIter(root))
	// 递归中序遍历
	fmt.Printf("递归中序遍历  	 =%v\n", InorderTraversal(root))
	// 非递归中序遍历
	fmt.Printf("非递归中序遍历 	 =%v\n", InorderTraversalIter(root))
	// 递归后序遍历
	fmt.Printf("递归后序遍历: 	 =%v\n\n", PostorderTraversal(root))
	// 非递归后序遍历
	fmt.Printf("非递归后序遍历:	 =%v\n", PostorderTraversalIter(root))
	// 层序遍历
	fmt.Printf("层序遍历:	 =%v\n", LevelOrder(root))
	// z字形层序遍历
	fmt.Printf("z字形层序遍历:	 =%v\n", ZigzagLevelOrder(root))

	single := singleton.GetSingleInstance()
	single.Value = 20
	single.Show()
}
