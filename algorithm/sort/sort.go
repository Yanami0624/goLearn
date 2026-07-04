package mysort

import (
	. "Gocodes/algorithm"
	"math/rand"
)

func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func RandomArr(n int) []MyInt {
	arr := make([]MyInt, n)
	for i := range n {
		l := rand.Intn(n)
		for arr[l] != 0 {
			l = (l + 1) % n
		}
		arr[l] = MyInt(i + 1)
	}
	return arr
}

func ComplexArr(n int) []MyInt {
	arr := make([]MyInt, n)
	for i := range n {
		arr[i] = MyInt(rand.Intn(n * 2))
	}
	return arr
}
