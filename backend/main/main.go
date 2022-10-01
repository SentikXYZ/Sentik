package main

import (
    "net/http"
	"log"
	"os/exec"
	"hackmit/mr"
    "github.com/gin-gonic/gin"
)

const (
	//Number of Mappers
	M			int = 5
)
type ticker struct {
	Ticker		string
	Name		string
	Date		string //UTC Format
}

func fetchTweets() {
	return ticker{
		Ticker: "VAPE",
		Name: "vape juice"
		Name: time.UTC().Now().String()
	}
}

func getTickers(c *gin.Context) {
	//Spawn map workers

	c.IndentedJSON(http.StatusOK, fetchTweets())
}

func main() {
	router := gin.Default()
    router.GET("/tickers", getTickers)

    router.Run("localhost:8080")
}



