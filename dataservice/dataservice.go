package dataservice

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func ParseRawData(rawData string) dataframe.DataFrame {
	df := dataframe.ReadCSV(strings.NewReader(rawData))
	return df
}

func ReadCSV(filePath string) dataframe.DataFrame {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	df := dataframe.ReadCSV(file)
	return df
}

func AddColumn(df *dataframe.DataFrame, values []float64, name string) dataframe.DataFrame {
	s := series.New(values, series.Float, name)
	// s.Append([]float64{1})
	return df.Mutate(s)
}

func WriteToFile(df dataframe.DataFrame, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(df)
	df.WriteCSV(file)
	defer file.Close()
}
