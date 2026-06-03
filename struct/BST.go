package data_struct

import "fmt"

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

func (root *BSTNode) Add(ele int) {
	if root == nil {
		root = new(BSTNode)
		root.val = ele
		return
	}
	cur := root
	for {
		switch {
		case ele > cur.val:
			if cur.right != nil {
				cur = cur.right
			} else {
				cur.right = new(BSTNode{val: ele})
			}
		case ele < cur.val:
			if cur.left != nil {
				cur = cur.right
			} else {
				cur.left = new(BSTNode{val: ele})
			}
		default:
			return
		}
	}
}

func BuildBST(n int) *BSTNode {
	root := BSTNode{val: 0}
	for i := 1; i < n; i++ {
		root.Add(i)
	}
	return &root
}

func (root *BSTNode) Print() {
	if root == nil {
		return
	}
	root.left.Print()
	fmt.Printf("%d ", root.val)
	root.right.Print()
}