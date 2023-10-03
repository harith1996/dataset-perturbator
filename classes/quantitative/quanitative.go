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

func ComputeDiff(q1 Quantitative, q2 Quantitative) float64 {
	return q1.value - q2.value
}

func GetSortedList(quantObjects []Quantitative) []Quantitative {
	sort.SliceStable(quantObjects, func(i, j int) bool {
		return ComputeDiff(quantObjects[i], quantObjects[j]) < 0
	})
	return quantObjects
}
