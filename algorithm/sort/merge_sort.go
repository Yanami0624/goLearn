package mysort

import . "Gocodes/algorithm"

func Sort_merge[T Comp](arr []T) {
	if len(arr) <= 1 {
		return
	}
	if len(arr) == 2 {
		if arr[0].Greater(arr[1]) {
			Swap(&arr[0], &arr[1])
		}
		return
	}

	r := len(arr)
	mid := (0 + r) / 2
	Sort_merge(arr[:mid])
	Sort_merge(arr[mid:])
	merge(arr)
}

func merge[T Comp](arr []T) {
	len := len(arr)
	mid := len / 2
	ret := make([]T, 0)
	pl, pr := 0, mid
	for pl != mid && pr != len {
		if arr[pl].Greater(arr[pr]) {
			ret = append(ret, arr[pr])
			pr++
		} else {
			ret = append(ret, arr[pl])
			pl++
		}
	}
	ret = append(ret, arr[pl:mid]...)
	ret = append(ret, arr[pr:]...)
	copy(arr, ret)
}