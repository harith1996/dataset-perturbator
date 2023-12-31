package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-gota/gota/series"
	"github.com/go-gota/gota/dataframe"

	"example.com/dp/calculators"
	"example.com/dp/dataservice"
	"example.com/dp/perturbators"
	"example.com/dp/structs"
	"github.com/google/uuid"
)

var pertRequests = []structs.PertRequest{
	{ID: uuid.NewString(), RawData: "x,y\n1,2\n2,3", Perturb: "addNoise", PerturbLevel: 1},
	{ID: uuid.NewString(), RawData: "x,y\n14,27\n23,38", Perturb: "jitter", PerturbLevel: 2},
	{ID: uuid.NewString(), RawData: "x,y\n19,22\n25,32", Perturb: "downsample", PerturbLevel: 3},
}

// Responds with the list of all perturbation requests as JSON.
func getPertRequests(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pertRequests)
}

// postPertRequest adds a pertRequest from JSON received in the request body.
func postPertRequest(c *gin.Context) {
	var newPertReq structs.PertRequest

	// Call BindJSON to bind the received JSON to
	// newPertRequest.
	if err := c.BindJSON(&newPertReq); err != nil {
		return
	}

	// Add an ID
	newPertReq.ID = uuid.NewString()

	// Add the new pertRequest to the slice.
	pertRequests = append(pertRequests, newPertReq)
	fmt.Println(perturbators.ApplyPert(newPertReq))
	c.IndentedJSON(http.StatusCreated, newPertReq)
}

func addTimeMapFields(c *gin.Context) {
	df := dataservice.ReadCSV("./data/monitoring_cruises_Feb2011_Dec2021_stationTF0286.csv")
	timeList := df.Col("Time").Records()
	customLayout := "02-01-2006 15:04:05"
	diffs := calculators.GetTimeMapLists(customLayout, timeList)
	diff0 := make([]float64, 0)
	diff1 := make([]float64, 0)
	for i, _ := range diffs {
		diff0 = append(diff0, diffs[i][0])
		diff1 = append(diff1, diffs[i][1])
	}
	df = dataservice.AddColumn(&df, diff0, "diffPrev")
	df = dataservice.AddColumn(&df, diff1, "diffNext")
	dataservice.WriteToFile(df, "timeMapAdded.csv")
	c.IndentedJSON(http.StatusOK, diffs)
}

func addQuantDiffFields(c *gin.Context) {
	df := dataservice.ReadCSV("./data/monitoring_cruises_Feb2011_Dec2021_stationTF0286.csv")
	fieldName := c.Query("fieldName")
	linearOrderBy := c.Query("linearOrderBy")
	sortedDf := df.Arrange(dataframe.Sort(linearOrderBy))
	rawRecords := sortedDf.Col(fieldName).Records()
	parsedRecords := make([]float64, len(rawRecords))
 
    for i, s := range rawRecords {
        num, err := strconv.ParseFloat(s, 64)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        parsedRecords[i] = num
    }
	diffs := calculators.GetQuantDiffs(parsedRecords)
	diff0 := make([]float64, 0)
	diff1 := make([]float64, 0)
	for i, _ := range diffs {
		diff0 = append(diff0, diffs[i][0])
		diff1 = append(diff1, diffs[i][1])
	}
	df = dataservice.AddColumn(&df, diff0, "diffPrev"+"_"+fieldName)
	df = dataservice.AddColumn(&df, diff1, "diffNext"+"_"+fieldName)
	dataservice.WriteToFile(df, "timeMapAdded.csv")
	c.IndentedJSON(http.StatusOK, diffs)
}


func addExpeditionField(c *gin.Context) {
	df_data := dataservice.ReadCSV("./timeMapAdded.csv")
	df_exped := dataservice.ReadCSV("./data/monitoring_cruises_Feb2011_Dec2021_stationTF0286_expedition_details.csv")
	layout_data := "02-01-2006 15:04:05"
	layout_exped := "02-01-06 15:04:05"
	exp_starts := df_exped.Col("Time (start)").Records()
	exp_ends := df_exped.Col("Time (end)").Records()

	for i, _ := range exp_ends {
		exp_starts[i] = calculators.AddEndOfDayTime(exp_starts[i])
		exp_ends[i] = calculators.AddEndOfDayTime(exp_ends[i])
	}

	data_timestamps := df_data.Col("Time").Records()
	exp_names := make([]string, 0)
	for _, t := range data_timestamps {
		for j := 0; j < len(exp_ends); j++ {
			if calculators.IsInTimeInterval(layout_data, layout_exped, t, []string{exp_starts[j], exp_ends[j]}) {
				exp_names = append(exp_names, df_exped.Col("Cruise Number").Records()[j])
			}
		}
	}
	df_data = df_data.Mutate(series.New(exp_names, series.String, "Expedition Number"))
	dataservice.WriteToFile(df_data, "expeditionAdded.csv")
	c.IndentedJSON(http.StatusOK, exp_names)
}

func main() {
	router := gin.Default()
	router.GET("/pertRequests", getPertRequests)
	router.POST("/pertRequests", postPertRequest)
	router.GET("/addTimeMapFields", addTimeMapFields)
	router.GET("/addExpeditionFields", addExpeditionField)
	router.Run("localhost:8080")
}
