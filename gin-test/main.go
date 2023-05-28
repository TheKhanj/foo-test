package main

import (
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "siktir")
}

func main() {
	os.Setenv("GOMAXPROCS", "1")
	runtime.GOMAXPROCS(1)

	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
