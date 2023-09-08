package perturbators

import (
	"fmt"
	"strings"
	"structs"

	"github.com/go-gota/gota/dataframe"
)

func ParseRawData(rawData string) dataframe.DataFrame {
	df := dataframe.ReadCSV(strings.NewReader(rawData))
	return df
}

func ApplyPert(request structs.PertRequest) dataframe.DataFrame {
	message := fmt.Sprintf("Applying %v to dataset", request.Perturb)
	fmt.Println(message)
	df := ParseRawData(request.RawData)
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
	xMean := xStats[1]
	yMean := yStats[1]
	xStdDev := xStats[2]
	yStdDev := yStats[2]
	xMin := xStats[3]
	yMin := yStats[3]
	fmt.Println(xMean, yMean, xStdDev, yStdDev, xMin, yMin)
	//add random noise to each column

	//return new dataframe
	return df, "noise added!"
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
