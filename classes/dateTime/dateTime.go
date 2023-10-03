package dateTime

import (
	"sort"
	"time"
)

func GetObject(layout string, timeString string) (time.Time, error) {
	return time.Parse(layout, timeString)
}

func AddEndOfDayTime(dateString string) string {
	return dateString + " 23:59:59"
}

func ComputeDiff(t1 time.Time, t2 time.Time) float64 {
	return t1.Sub(t2).Abs().Seconds()
}

func GetSortedList(timeObjects []time.Time) []time.Time {
	sort.SliceStable(timeObjects, func(i, j int) bool {
		return timeObjects[i].Before(timeObjects[j])
	})
	return timeObjects
}


// checks if a given "timeString" is within a time interval
func IsInTimeInterval(layoutT string, layoutI string, timeString string, intervalStrings []string) bool {
	timeObject, _ := GetObject(layoutT, timeString)
	time0, _ := GetObject(layoutI, intervalStrings[0])
	time1, _ := GetObject(layoutI, intervalStrings[1])
	return time0.Before(timeObject) && timeObject.Before(time1)
}


func GetTimeObjects(layout string, times []string) []time.Time {
	timeObjects := []time.Time{}
	for _, timeString := range times {
		timeObject, _ := GetObject(layout, timeString)
		timeObjects = append(timeObjects, timeObject)
	}
	sort.SliceStable(timeObjects, func(i, j int) bool {
		return timeObjects[i].Before(timeObjects[j])
	})
	return timeObjects
}

