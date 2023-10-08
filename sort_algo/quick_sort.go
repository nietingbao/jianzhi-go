package sort_algo

func SortQ(a []int) {
	if len(a) <= 1 {
		return
	}
	quickS(a, 0, len(a)-1)
}

func quickS(a []int, l, r int) {
	if l >= r {
		return
	}
	p, i, j := a[l], l+1, r
	for {
		for a[i] <= p && i < r {
			i++
		}
		for a[j] >= p && j > l {
			j--
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	if j >= l {
		a[j], a[l] = a[l], a[j]
	}

	quickS(a, l, j-1)
	quickS(a, j+1, r)
}
