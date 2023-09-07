package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PertRequest struct {
	RawData      string `json:"rawData"`
	Perturb      string `json:"perturb"`      // "add_noise", "downsample", "jitter"
	PerturbLevel int    `json:"perturbLevel"` // 1-3
}

var pertRequests = []PertRequest{
	{RawData: "x,y\n1,2\n2,3", Perturb: "add_noise", PerturbLevel: 1},
	{RawData: "x,y\n14,27\n23,38", Perturb: "jitter", PerturbLevel: 2},
	{RawData: "x,y\n19,22\n25,32", Perturb: "downsample", PerturbLevel: 3},
}

// getPertRequests responds with the list of all albums as JSON.
func getPertRequests(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pertRequests)
}

func main() {
	router := gin.Default()
	router.GET("/pertRequests", getPertRequests)
	router.Run("localhost:8080")
}
