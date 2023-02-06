package main

import (
	"log"
	"net/http"

	"github.com/ericthomasca/go-sqlite/server/models"
	"github.com/gin-gonic/gin"
)

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

	r := gin.Default()

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("book", getBooks)
		v1.GET("book/:id", getBooksById)
		v1.POST("book", addBook)
		v1.PUT("book/:id", updateBook)
		v1.DELETE("book/:id", deleteBook)
		v1.OPTIONS("book", options)
	}

	r.Run()
}

func getBooks(c *gin.Context) {
	books, err := models.GetBooks(5)
	checkErr(err)

	if books == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": books})
	}
}

func getBooksById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "getBooksById " + id + " called"})
}

func addBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "addBook called"})
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "updateBook " + id + " called"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "deleteBook " + id + " called"})
}

func options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "options called"})
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}