package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"perturbators"

	"structs"
)

var pertRequests = []structs.PertRequest{
	{RawData: "x,y\n1,2\n2,3", Perturb: "add_noise", PerturbLevel: 1},
	{RawData: "x,y\n14,27\n23,38", Perturb: "jitter", PerturbLevel: 2},
	{RawData: "x,y\n19,22\n25,32", Perturb: "downsample", PerturbLevel: 3},
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
