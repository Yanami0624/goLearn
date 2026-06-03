package main

import (
	. "Gocodes/algorithm/binary_search"
	. "Gocodes/algorithm/sort"
	"fmt"
)

func main() {
	arr := ComplexArr(16)
	Sort_quick(arr)
	fmt.Println(arr)
	fmt.Println(BS_rightbound(arr, 17))
}
