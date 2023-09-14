package perturbators

import (
	"fmt"
	"strings"
	"example.com/dp/structs"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"gonum.org/v1/gonum/spatial/r1"
	"gonum.org/v1/gonum/stat/distmv"
)

func ParseRawData(rawData string) dataframe.DataFrame {
	df := dataframe.ReadCSV(strings.NewReader(rawData))
	return df
}

func ApplyPert(request structs.PertRequest) dataframe.DataFrame {
	message := fmt.Sprintf("Applying %v to dataset", request.Perturb)
	fmt.Println(message)
	df := ParseRawData(request.RawData)
	print(df.Records())
	method := AddNoise
	switch request.Perturb {
	case "addNoise":
		method = AddNoise
	case "jitter":
		method = Jitter
	case "downsample":
		method = Downsample
	}
	newDf, status := method(df, request.PerturbLevel)
	fmt.Println(status)
	return newDf
}

// adds random noise to the dataset, and returns the new dataset
// level is the amount of noise to add
func AddNoise(df dataframe.DataFrame, level int) (dataframe.DataFrame, string) {
	//get mean and std dev of each column
	xStats := df.Describe().Col("x").Float()
	yStats := df.Describe().Col("y").Float()
	// xMean := xStats[1]
	// yMean := yStats[1]
	// xStdDev := xStats[2]
	// yStdDev := yStats[2]
	xMin := xStats[3]
	yMin := yStats[3]
	xMax := xStats[7]
	yMax := yStats[7]
	xBound := r1.Interval{Min: xMin, Max: xMax}
	yBound := r1.Interval{Min: yMin, Max: yMax}

	//create a multivariate uniform distribution
	dist := distmv.NewUniform([]r1.Interval{xBound, yBound}, nil)

	//concatenate noise to df
	xSeries := df.Col("x")
	ySeries := df.Col("y")
	//draw samples from the distribution
	for i := 0; i < level*len(df.Records())/4; i++ {
		sample := []float64{0, 0}
		dist.Rand(sample)
		xSeries.Append(series.Floats(sample[0]))
		ySeries.Append(series.Floats(sample[1]))
	}
	newDf := dataframe.New(xSeries, ySeries)
	//return new dataframe
	return newDf, "noise added!"
}

func Jitter(df dataframe.DataFrame, level int) (dataframe.DataFrame, string) {
	// Return a greeting that embeds the name in a message.
	fmt.Println("Hi! I am the jitterer :)")
	return df, "points jittered!"
}

func Downsample(df dataframe.DataFrame, level int) (dataframe.DataFrame, string) {
	// Return a greeting that embeds the name in a message.
	fmt.Println("Hi! I am the downsampler :)")
	return df, "downsampled!"
}
