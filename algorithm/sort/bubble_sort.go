package mysort

import . "Gocodes/algorithm"

func Sort_bubble[T Comp](arr []T) {
	for i := range arr {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j].Greater(arr[j+1]) {
				Swap(&arr[j], &arr[j+1])
			}
		}
	}
}
