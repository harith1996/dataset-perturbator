package calculators

import (
	"example.com/dp/classes/dateTime"
	"example.com/dp/classes"
	"fmt"
)
// supply a list of values and a diff function
// returns 2 values per list item - diffPrev and diffNext
func GetTimeMapLists(layout string, times []string, tuple classes.DateTimeTuple) [][]float64 {
	fmt.Println("Computing time map lists")
	// layout := "2006-01-02T15:04:05Z"
	// customLayout := "02-01-2006 15:04:05"
	timeObjects := dateTime.GetTimeObjects(layout, times)

	diffBefore := []float64{}
	diffAfter := []float64{}

	for i, timeObject := range timeObjects {
		if i == 0 {
			diffBefore = append(diffBefore, 0)
			diffAfter = append(diffAfter, dateTime.ComputeDiff(timeObjects[i+1], timeObject))
		} else if i == len(timeObjects)-1 {
			diffAfter = append(diffAfter, 0)
			diffBefore = append(diffBefore, dateTime.ComputeDiff(timeObject, timeObjects[i-1]))
		} else {
			diffBefore = append(diffBefore, dateTime.ComputeDiff(timeObject, timeObjects[i-1]))
			diffAfter = append(diffAfter, dateTime.ComputeDiff(timeObjects[i+1], timeObject))
		}
	}

	//return the two lists as a single list
	diffs := make([][]float64, 0)
	for i := range diffBefore {
		diffs = append(diffs, []float64{diffBefore[i], diffAfter[i]})
	}
	return diffs
}

// takes in a list of sorted measurements, and returns the difference
// between adjacent measurements for each value
func ComputeMeasurementDiffs(sortedValues []float64) {

}
