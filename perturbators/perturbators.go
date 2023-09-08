package perturbators

import (
	"fmt"
	"strings"
	"structs"
	"github.com/go-gota/gota/dataframe"
)

func ParseRawData(rawData string) dataframe.DataFrame {
	df := dataframe.ReadCSV(strings.NewReader(rawData))
	fmt.Println(df)
	return df
}

func ApplyPert(request structs.PertRequest) string {
	message := fmt.Sprintf("Applying %v to dataset", request.Perturb)
	ParseRawData(request.RawData)
	return message
}

// AddNoise returns a greeting for the named person.
func AddNoise(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome! I am the noise adder :)", name)
	return message
}

func Jitter(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome! I am the jitterer :)", name)
	return message
}

func Downsample(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome! I am the downsampler :)", name)
	return message
}
