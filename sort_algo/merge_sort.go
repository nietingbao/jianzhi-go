package sort_algo

func MergeSort(a []int) {
	mergeS(a, 0, len(a)-1)
}

func mergeS(a []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergeS(a, l, mid)
	mergeS(a, mid+1, r)
	mergeSorted(a, l, mid, r)
}

func mergeSorted(a []int, l, m, r int) {
	tmp := make([]int, 0, r-l+1)
	low, high := l, m+1
	for low <= m && high <= r {
		if a[low] < a[high] {
			tmp = append(tmp, a[low])
			low++
		} else {
			tmp = append(tmp, a[high])
			high++
		}
	}
	for low <= m {
		tmp = append(tmp, a[low])
		low++
	}
	for high <= r {
		tmp = append(tmp, a[high])
		high++
	}
	for i := 0; i < len(tmp); i++ {
		a[l+i] = tmp[i]
	}
}
