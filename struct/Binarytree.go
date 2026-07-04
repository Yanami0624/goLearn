package data_struct

import "fmt"

type BinaryTreeNode struct {
	val   int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func BuildBinaryTree(n int) *BinaryTreeNode {
	if n <= 0 {
		return nil
	}
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i
	}
	root := BinaryTreeNode{
		val: arr[0],
	}
	q := Queue[*BinaryTreeNode]{}
	q.Push(&root)

	for i := 1; i < n; i++ {
		n := new(BinaryTreeNode)
		n.val = arr[i]

		cur := q.Front()
		if cur.left == nil {
			cur.left = n
		} else {
			cur.right = n
			q.Pop()
		}
		q.Push(n)
	}
	return &root
}

func (root *BinaryTreeNode) Print() {
	f := func(n *BinaryTreeNode) (_ any) {
		fmt.Printf("%d ", n.val)
		return
	}
	root.Travel_PreOrder(f)
}

func (root *BinaryTreeNode) Travel_PreOrder(f func(*BinaryTreeNode) any) {
	if root == nil {
		return
	}
	f(root)
	if root.left != nil {
		root.left.Travel_PreOrder(f)
	}
	if root.right != nil {
		root.right.Travel_PreOrder(f)
	}
}

func (root *BinaryTreeNode) Travel_InOrder(f func(*BinaryTreeNode) any) {
	if root == nil {
		return
	}
	if root.left != nil {
		root.left.Travel_PreOrder(f)
	}
	f(root)
	if root.right != nil {
		root.right.Travel_PreOrder(f)
	}
}

func (root *BinaryTreeNode) Travel_PostOrder(f func(*BinaryTreeNode) any) {
	if root == nil {
		return
	}
	if root.left != nil {
		root.left.Travel_PreOrder(f)
	}
	if root.right != nil {
		root.right.Travel_PreOrder(f)
	}
	f(root)
}
