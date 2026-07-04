package binarysearch

import (
	. "Gocodes/algorithm"
)

func BS_exact[T Comp](arr []T, target T) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r) / 2
		switch {
		case target.Greater(arr[mid]):
			l = mid + 1
		case arr[mid].Greater(target):
			r = mid - 1
		default:
			return mid
		}
	}
	if arr[r].Equal(target) {
		return r
	}
	return -1
}

func BS_leftbound[T Comp](arr []T, target T) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r) / 2
		switch {
		case target.Greater(arr[mid]):
			l = mid + 1
		case arr[mid].Greater(target):
			r = mid - 1
		default:
			r = mid - 1
		}
	}
	return l
}

func BS_rightbound[T Comp](arr []T, target T) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r) / 2
		switch {
		case target.Greater(arr[mid]):
			l = mid + 1
		case arr[mid].Greater(target):
			r = mid - 1
		default:
			l = mid + 1
		}
	}
	return r
}
