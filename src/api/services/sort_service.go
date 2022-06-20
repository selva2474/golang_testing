package services

import (
	"github.com/selva2474/golang_testing/src/api/utils/sort"
)

func Sort(elements []int) {
	if len(elements) <= 20000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
