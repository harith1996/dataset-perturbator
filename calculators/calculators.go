package calculators

import (
	"fmt"
	"sort"
	"time"
)

func getTimeObjects(layout string, times []string) []time.Time {
	timeObjects := []time.Time{}
	for _, timeString := range times {
		timeObject, _ := time.Parse(layout, timeString)
		timeObjects = append(timeObjects, timeObject)
	}
	sort.SliceStable(timeObjects, func(i, j int) bool {
		return timeObjects[i].Before(timeObjects[j])
	})
	return timeObjects
}

func GetTimeMapLists(layout string, times []string) [][]float64 {
	fmt.Println("Computing time map lists")
	// layout := "2006-01-02T15:04:05Z"
	// customLayout := "02-01-2006 15:04:05"
	timeObjects := getTimeObjects(layout, times)

	diffBefore := []float64{}
	diffAfter := []float64{}

	for i, timeObject := range timeObjects {
		if i == 0 {
			diffBefore = append(diffBefore, 0)
			diffAfter = append(diffAfter, timeObjects[i+1].Sub(timeObject).Abs().Seconds())
		} else if i == len(timeObjects)-1 {
			diffAfter = append(diffAfter, 0)
			diffBefore = append(diffBefore, timeObject.Sub(timeObjects[i-1]).Abs().Seconds())
		} else {
			diffBefore = append(diffBefore, timeObject.Sub(timeObjects[i-1]).Abs().Seconds())
			diffAfter = append(diffAfter, timeObjects[i+1].Sub(timeObject).Abs().Seconds())
		}
	}

	//return the two lists as a single list
	diffs := make([][]float64, 0)
	for i := range diffBefore {
		diffs = append(diffs, []float64{diffBefore[i], diffAfter[i]})
	}
	return diffs
}

func GetQuantDiffs(values []float64) [][]float64 {
	diffBefore := []float64{}
	diffAfter := []float64{}

	for i, v := range values {
		if i == 0 {
			diffBefore = append(diffBefore, 0)
			diffAfter = append(diffAfter, values[i+1]-v)
		} else if i == len(values)-1 {
			diffAfter = append(diffAfter, 0)
			diffBefore = append(diffBefore, v-values[i-1])
		} else {
			diffBefore = append(diffBefore, v-values[i-1])
			diffAfter = append(diffAfter, values[i+1]-v)
		}
	}

	//return the two lists as a single list
	diffs := make([][]float64, 0)
	for i := range diffBefore {
		diffs = append(diffs, []float64{diffBefore[i], diffAfter[i]})
	}
	return diffs
}

// checks if a given "timeString" is within a time interval
func IsInTimeInterval(layoutT string, layoutI string, timeString string, intervalStrings []string) bool {
	timeObject, _ := time.Parse(layoutT, timeString)
	time0, _ := time.Parse(layoutI, intervalStrings[0])
	time1, _ := time.Parse(layoutI, intervalStrings[1])
	return time0.Before(timeObject) && timeObject.Before(time1)
}

func AddEndOfDayTime(dateString string) string {
	return dateString + " 23:59:59"
}

// takes in a list of sorted measurements, and returns the difference
// between adjacent measurements for each value
func ComputeMeasurementDiffs(sortedValues []float64) {

}
