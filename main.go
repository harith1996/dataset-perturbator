package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

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

func main() {
	router := gin.Default()
	router.GET("/pertRequests", getPertRequests)
	router.POST("/pertRequests", postPertRequest)
	router.Run("localhost:8080")
}
