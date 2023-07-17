package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/wangjohn/quickselect"
	"math/rand"
	"sort"
	"testing"
)

const (
	size = 2_000_000
	k    = 100_000
)

var slice []float64

func init() {
	slice = make([]float64, size)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Float64()
	}
}

func TestQuickSelect(t *testing.T) {
	limit := 1000
	arr1 := make([]int, 10_000)
	arr2 := make([]int, 10_000)

	for i := 0; i < len(arr1); i++ {
		arr1[i] = rand.Int()
		arr2[i] = arr1[i]
	}

	sort.Ints(arr1)

	m := make(map[int]struct{}, 0)
	for i := 0; i < limit; i++ {
		m[arr1[i]] = struct{}{}
	}

	quickselect.QuickSelect(quickselect.IntSlice(arr2), limit)
	arr2 = arr2[:limit]

	for _, v := range arr2 {
		_, ok := m[v]
		assert.Equal(t, true, ok)
	}
}

func TestQuickSelectOwn(t *testing.T) {
	limit := 5000
	arr1 := make([]float64, 10_000)
	arr2 := make([]float64, 10_000)

	for i := 0; i < len(arr1); i++ {
		arr1[i] = rand.Float64()
		arr2[i] = arr1[i]
	}

	sort.Slice(arr1, func(i, j int) bool { return arr1[i] < arr1[j] })

	m := make(map[float64]struct{}, 0)
	for i := 0; i < limit; i++ {
		m[arr1[i]] = struct{}{}
	}

	QuickSelect(arr2, 0, len(arr2)-1, limit)

	arr2 = arr2[:limit]

	for _, v := range arr2 {

		_, ok := m[v]
		assert.Equal(t, true, ok)
	}
}

func BenchmarkSearchTopkWithSorting(b *testing.B) {
	arr := make([]float64, len(slice))
	for i := 0; i < len(slice); i++ {
		arr[i] = slice[i]
	}
	sorted := make([]int, len(arr))
	for i := 0; i < len(sorted); i++ {
		sorted[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Slice(sorted, func(i, j int) bool {
			return arr[sorted[i]] > arr[sorted[j]]
		})
	}
}

func BenchmarkSearchTopkWithQuickSelect(b *testing.B) {
	arr := make([]float64, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = float64(slice[i])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quickselect.QuickSelect(quickselect.Float64Slice(arr), k)
	}
}

func BenchmarkSearchTopkWithOwnQuickSelect(b *testing.B) {
	arr := make([]float64, size)

	for i := 0; i < len(arr); i++ {
		arr[i] = float64(slice[i])
		//arr[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//quickselect.QuickSelect(quickselect.Float64Slice(arr), k)
		QuickSelect(arr, 0, len(arr)-1, len(arr)-k)
	}
}
