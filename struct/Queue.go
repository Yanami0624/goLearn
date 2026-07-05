package data_struct

import "fmt"

type Queue[T any] struct {
	arr []T
}

func (q *Queue[T]) Push(ele T) {
	q.arr = append(q.arr, ele)
}

func (q *Queue[T]) Pop() T {
	ret := q.arr[0]
	q.arr = q.arr[1:]
	return ret
}

func (q *Queue[T]) Front() T {
	return q.arr[0]
}

func (q *Queue[T]) Size() int {
	return len(q.arr)
}

func (q *Queue[T]) Print() {
	fmt.Println(q.arr)
}
