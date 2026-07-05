package data_struct

type TreeNode[T any] struct {
	val  T
	sons []*TreeNode[T]
}
