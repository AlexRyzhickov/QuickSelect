package main

import "fmt"

func Partition(v []float64, left int, right int) int {

	pivot := v[(right+left)/2]
	l := left
	r := right

	for l <= r {
		for v[l] < pivot {
			l++
		}
		for v[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		v[l], v[r] = v[r], v[l]
		r--
		l++
	}

	return r
}

func QuickSelect(v []float64, left int, right int, k int) {
	if right == left {
		return
	}

	mid := Partition(v, left, right)
	if k > mid {
		QuickSelect(v, mid+1, right, k)
	} else {
		QuickSelect(v, left, mid, k)
	}
}

func main() {
	v := []float64{9, 3, 6, 1, 8, 2, 5, 7, 4}
	k := 3

	QuickSelect(v, 0, len(v)-1, len(v)-k)

	fmt.Println(v[len(v)-k:])
	//for i := len(v) - k; i < len(v); i++ {
	//	fmt.Print(v[i], " ")
	//}
}
