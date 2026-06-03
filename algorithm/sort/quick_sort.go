package mysort

import . "Gocodes/algorithm"

func Sort_quick[T Comp](arr []T) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[len(arr)-1]
	ptr := 0
	for i := 0; i < len(arr)-1; i++ {
		if pivot.Greater(arr[i]) {
			Swap(&arr[i], &arr[ptr])
			ptr++
		}
	}
	Swap(&arr[len(arr)-1], &arr[ptr])

	Sort_quick(arr[:ptr])
	Sort_quick(arr[ptr+1:])
}