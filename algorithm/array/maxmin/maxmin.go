package maxmin

/*

给定数组a1,a2,a3...an，要求找出数组中的最大值和最小值。假设数组中的值两两各不相同。

*/

func MiuerImpl(arr []int, l, r int) (max, min int) {
	if arr == nil {
		return
	}

	if r-l == 0 {
		return arr[l], arr[r]
	}

	if r-l == 1 {
		if arr[l] > arr[r] {
			return arr[l], arr[r]
		}
		return arr[r], arr[l]
	}

	mid := (l + r) / 2

	lmax, lmin := MiuerImpl(arr, l, mid)
	rmax, rmin := MiuerImpl(arr, mid+1, r)

	if lmax > rmax {
		max = lmax
	} else {
		max = rmax
	}

	if lmin > rmin {
		min = rmin
	} else {
		min = lmin
	}

	return
}
