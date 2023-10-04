package quantitative

import (
	"sort"
)

type Quantitative struct {
	value   float64
}

func GetObject(t Quantitative) Quantitative {
	return t
}


func GetSortedList(quantObjects []Quantitative) []Quantitative {
	sort.SliceStable(quantObjects, func(i, j int) bool {
		return ComputeDiff(quantObjects[i], quantObjects[j]) < 0
	})
	return quantObjects
}
