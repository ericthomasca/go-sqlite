package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("book", getBooks)
		v1.GET("book/:id", getBooksById)
		v1.POST("book", addBook)
		v1.PUT("book/:id", updateBook)
		v1.DELETE("person/:id", deleteBook)
		v1.OPTIONS("book", options)
	}

	r.Run()
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getBooks called"})
}

func getBooksById(c *gin.Context) {
	id := c.Param("Id")
	c.JSON(http.StatusOK, gin.H{"message": "getBooksById " + id + " called"})
}

func addBook(c *gin.Context) {

}
