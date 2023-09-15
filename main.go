package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

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

func getTimeMapFields(c *gin.Context) {
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

func main() {
	router := gin.Default()
	router.GET("/pertRequests", getPertRequests)
	router.POST("/pertRequests", postPertRequest)
	router.GET("/getTimeMapFields", getTimeMapFields)
	router.Run("localhost:8080")
}
