package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// adding few album
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Contrate", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbumById)
	router.POST("/albums", postAlbum)
	router.Run("localhost:9999")
}

// writing the endpoints
// writing a handler to return all items

// getAlbums function creates JSON from the slices of album struct, writing the json in response
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums to add a new Album from a json received in the request body
func postAlbum(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add an album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumById to extract the ID in the request Path
func getAlbumById(c *gin.Context) {
	//get the id from context
	id := c.Param("id")

	// we will loop over albums to get the id matching the request id and return it as json
	for _, record := range albums {
		if record.ID == id {
			c.IndentedJSON(http.StatusOK, record)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
