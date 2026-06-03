package data_struct

import "fmt"

type Stack[T any] struct {
	arr []T
}

func (s *Stack[T]) Push(ele T) {
	s.arr = append(s.arr, ele)
}

func (s *Stack[T]) Pop() T {
	ret := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-2]
	return ret
}

func (s *Stack[T]) Top() T {
	return s.arr[len(s.arr)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.arr)
}

func (s *Stack[T]) Print() {
	fmt.Println(s.arr)
}