package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initTree() *TreeNode {
	t := &TreeNode{Val: 1}
	// t.Left = &TreeNode{Val: 2}
	// t.Right = &TreeNode{Val: 2}
	// t.Left.Left = &TreeNode{Val: 3}
	// t.Left.Right = &TreeNode{Val: 4}
	// t.Right.Left = &TreeNode{Val: 4}
	// t.Right.Right = &TreeNode{Val: 3}
	t.Right = &TreeNode{Val: 2}
	t.Right.Right = &TreeNode{Val: 3}
	t.Right.Right.Right = &TreeNode{Val: 4}
	t.Right.Right.Right.Right = &TreeNode{Val: 5}

	return t
}

func main() {
	t := initTree()
	res := hasPathSum1(t, 15)

	fmt.Println(res)
}

// 112/tree/path-sum
// post order traversal
func hasPathSum(root *TreeNode, targetSum int) bool {
	curSum := 0
	p := []*TreeNode{}
	node, pre := root, root
	for len(p) > 0 || node != nil {
		for node != nil {
			p = append(p, node)
			if curSum+node.Val == targetSum && node.Left == nil && node.Right == nil {
				return true
			} else {
				curSum += node.Val
			}
			node = node.Left

		}
		node = p[len(p)-1]
		p = p[:len(p)-1]
		// 右子树未遍历
		if node.Right != nil && node.Right != pre {
			p = append(p, node)
			node = node.Right
		} else { // 右子树不存在或者已遍历
			pre = node
			curSum -= node.Val
			node = nil
		}

	}

	return false
}

// 112/tree/path-sum
// recursion
func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum1(root.Left, targetSum-root.Val) || hasPathSum1(root.Right, targetSum-root.Val)
}

// 112/tree/path-sum
// level order traversal
func hasPathSum2(root *TreeNode, targetSum int) bool {

	return false
}

// 107/tree/binary-tree-level-order-traversal-ii
// level order traversal
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	p := []*TreeNode{}
	p = append(p, root)
	res = append(res, []int{root.Val})
	for i := 0; len(p) > 0; i++ {
		q := []*TreeNode{}
		arr := []int{}
		for j := 0; j < len(p); j++ {
			node := p[j]
			if node.Left != nil {
				q = append(q, node.Left)
				arr = append(arr, node.Left.Val)
			}
			if node.Right != nil {
				q = append(q, node.Right)
				arr = append(arr, node.Right.Val)
			}
		}
		if len(arr) > 0 {
			res = append(res, arr)
		}
		p = q
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}

// 110/tree/balanced-binary-tree
// recursion
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalance_abs(isBalance_height(root.Left)-isBalance_height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
func isBalance_height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return isBalance_max(isBalance_height(root.Left), isBalance_height(root.Right)) + 1
}
func isBalance_max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func isBalance_abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// 257/tree/binary-tree-paths
// todo
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	nodeArr := []*TreeNode{}
	arr := []*TreeNode{}
	arr = append(arr, root)
	for i := 0; len(arr) > 0; i++ {
		p := []*TreeNode{}
		for j := 0; j < len(arr); j++ {
			node := p[0]
			p = p[:len(p)-1]
			if node.Left == nil && node.Right == nil {
				nodeArr = append(nodeArr, node)
			}
			if node.Left != nil {
				p = append(p, node.Left)

			}
		}

	}

	return nil
}

// 226/tree/invert-binary-tree
// recursion
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 226/tree/invert-binary-tree
// iteration
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := []*TreeNode{}
	node := root
	queue = append(queue, node)
	for i := 0; len(queue) > 0; i++ {
		node = queue[0]
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}

// 104/tree/maximum-depth-of-binary-tree
// DFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	q := []*TreeNode{}
	q = append(q, root)
	for ; len(q) > 0; max++ {
		p := []*TreeNode{}
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return max
}

// 104/tree/maximum-depth-of-binary-tree
// recursion
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left < right {
		return right + 1
	}
	return left + 1
}

// 101/tree/symmetric-tree
// 迭代
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	q := []*TreeNode{}
	q = append(q, root)
	for i := 0; len(q) > 0; i++ {
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			if q[j] != nil {
				p = append(p, q[j].Left)
				p = append(p, q[j].Right)
			}
		}
		if len(p)%2 == 0 {
			head, tail := 0, len(p)-1
			for head < tail {
				if p[head] == nil && p[tail] == nil {
					head++
					tail--
					continue
				}
				if p[head] == nil || p[tail] == nil {
					return false
				}
				if p[head].Val != p[tail].Val {
					return false
				}
				head++
				tail--
			}
		} else if len(p) > 1 {
			return false
		}
		q = p
	}
	return true
}

// 101/tree/symmetric-tree
// 递归方式
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSymmetric1_check(root.Left, root.Right)
}
func isSymmetric1_check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetric1_check(left.Left, right.Right) && isSymmetric1_check(left.Right, right.Left)
}

// 101/tree/symmetric-tree
// 迭代官方解法
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return false
	}
	arr := []*TreeNode{}
	arr = append(arr, root.Left, root.Right)
	for len(arr) > 0 {
		p, q := arr[0], arr[1]
		arr = arr[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		arr = append(arr, p.Left)
		arr = append(arr, q.Right)

		arr = append(arr, p.Right)
		arr = append(arr, q.Left)
	}

	return true
}

// 100/tree/same_tree
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	pStack, qStack := []*TreeNode{}, []*TreeNode{}
	pStack = append(pStack, p)
	qStack = append(qStack, q)
	for len(qStack) > 0 && len(pStack) > 0 {
		p = pStack[len(pStack)-1]
		pStack = pStack[:len(pStack)-1]
		q = qStack[len(qStack)-1]
		qStack = qStack[:len(qStack)-1]
		if p.Val != q.Val {
			return false
		}
		if p.Left != nil && q.Left != nil {
			pStack = append(pStack, p.Left)
			qStack = append(qStack, q.Left)
		} else if (p.Left != nil && q.Left == nil) || (p.Left == nil && q.Left != nil) {
			return false
		}
		if p.Right != nil && q.Right != nil {
			pStack = append(pStack, p.Right)
			qStack = append(qStack, q.Right)
		} else if (p.Right != nil && q.Right == nil) || (p.Right == nil && q.Right != nil) {
			return false
		}

	}
	return true
}

func do(node *TreeNode) {
	fmt.Println(node.Val)
}

func PreOrderTraversal(t *TreeNode) {
	stack := []*TreeNode{}
	stack = append(stack, t)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		do(node)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

func InOrderTraversal(t *TreeNode) {
	stack := []*TreeNode{}
	node := t
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		do(node)
		node = node.Right
	}
}

func PostOrderTraversal(t *TreeNode) {
	type nodeExt struct {
		v    int
		node *TreeNode
	}
	stack := []*nodeExt{}
	curNode := &nodeExt{v: 0, node: t}
	stack = append(stack, curNode)
	for len(stack) > 0 {
		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curNode.v == 0 {
			curNode.v += 1
			stack = append(stack, curNode)
			if curNode.node.Right != nil {
				stack = append(stack, &nodeExt{v: 0, node: curNode.node.Right})
			}
			if curNode.node.Left != nil {
				stack = append(stack, &nodeExt{v: 0, node: curNode.node.Left})
			}
		} else {
			do(curNode.node)
		}
	}
}
func PostOrderTraversal1(t *TreeNode) {
	stack := []*TreeNode{}
	node := t
	pre := node
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 右子树不存在或者已遍历
		if node.Right == nil || node.Right == pre {
			do(node)
			pre = node
			node = nil
		} else { // 右子树未遍历
			stack = append(stack, node)
			node = node.Right
		}

	}
}

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	type nodeExt struct {
		dep  int
		node *TreeNode
	}
	queue := []*nodeExt{}
	res1 := []*nodeExt{}
	node := &nodeExt{dep: 0, node: root}
	queue = append(queue, node)
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		res1 = append(res1, node)
		if node.node.Left != nil {
			queue = append(queue, &nodeExt{dep: node.dep + 1, node: node.node.Left})
		}
		if node.node.Right != nil {
			queue = append(queue, &nodeExt{dep: node.dep + 1, node: node.node.Right})
		}
	}
	res := [][]int{}
	curDep := 0
	r := []int{}
	for _, v := range res1 {
		if v.dep == curDep {
			r = append(r, v.node.Val)
		} else {
			res = append(res, r)
			r = []int{}
			r = append(r, v.node.Val)
			curDep++
		}
	}
	res = append(res, r)
	return res
}

func LevelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	type ext struct {
		dep  int
		node *TreeNode
	}
	res := [][]int{}
	queue := []*ext{}
	curNode := &ext{dep: 0, node: root}
	queue = append(queue, curNode)
	curDep := 0
	arr := []int{}
	for len(queue) > 0 {
		curNode = queue[0]
		queue = queue[1:]
		if curNode.dep == curDep {
			arr = append(arr, curNode.node.Val)
		} else {
			res = append(res, arr)
			arr = []int{}
			arr = append(arr, curNode.node.Val)
			curDep++
		}
		if curNode.node.Left != nil {
			queue = append(queue, &ext{dep: curNode.dep + 1, node: curNode.node.Left})
		}
		if curNode.node.Right != nil {
			queue = append(queue, &ext{dep: curNode.dep + 1, node: curNode.node.Right})
		}
	}
	res = append(res, arr)
	return res
}
func LevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{}
	res := [][]int{}
	node := root
	q = append(q, node)
	for i := 0; len(q) > 0; i++ {
		p := []*TreeNode{}
		res = append(res, []int{})
		for j := 0; j < len(q); j++ {
			node = q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return res
}
