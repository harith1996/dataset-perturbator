package classes

import ("time")

type DataItem interface {
	ComputeDiff() float64
}

type DateTimeTuple struct {
	first time.Time
	second time.Time
}

type QuantTuple struct {
	first float64
	second float64
}

func (d DateTimeTuple) ComputeDiff() float64 {
	return d.first.Sub(d.second).Abs().Seconds()
}

func (q QuantTuple) ComputeDiff() float64 {
	return q.first - q.second
}

func GetTuple(typeName string) int64 {
	switch(typeName) {
	case "DateTime" : return DateTimeTuple
	}
}
