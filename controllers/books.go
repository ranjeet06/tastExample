package controllers

import (
	"example.com/testProject/newModule"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testProject/models"
)

// GET /books
// Get all books

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

func FindBooks(c *gin.Context) {

	h := testHeader{}
	err := c.BindHeader(&h)
	if err != nil {
		log.Panicln(err)
	}

	c.Header("Age", "0130")

	c.SetCookie("test_cookie", "test_value", 10, "/", "localhost:8080", false, true)

	if c.ContentType() == "application/xml" {
		c.XML(http.StatusOK, gin.H{"data": models.DB})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": models.DB})
	}
	newModule.Hello()
	c.JSON(200, gin.H{"Rate": h.Rate, "Domain": h.Domain})

}

//GET /books/:id
// Find a book

func FindBook(c *gin.Context) {
	id := c.Param("id")

	for _, a := range models.DB {
		if a.ID == id {
			if c.ContentType() == "application/xml" {
				c.XML(http.StatusOK, gin.H{"data": a})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": a})
			}

			return
		}
	}

	if c.ContentType() == "application/xml" {
		c.XML(http.StatusBadRequest, gin.H{"data": "Book not found"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Book not found"})
	}

}

// POST/book
// create new book

func CreateBook(c *gin.Context) {

	var input models.Book

	if c.ContentType() == "application/xml" {
		if err := c.BindXML(&input); err != nil {
			log.Panicln(err)
		}
	} else {
		if err := c.BindJSON(&input); err != nil {
			log.Panicln(err)
		}
	}

	for _, a := range models.DB {
		if a.ID == input.ID {

			if c.ContentType() == "application/xml" {
				c.XML(http.StatusOK, gin.H{"data": "Book ID already exist"})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": "Book ID already exist"})
			}
			return
		}
	}
	models.DB = append(models.DB, input)

	if c.ContentType() == "application/xml" {
		c.XML(http.StatusOK, gin.H{"data": models.DB})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": models.DB})
	}

}

// PATCH /books/:id
// Update a book

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var input models.Book

	if c.ContentType() == "application/xml" {
		if err := c.BindXML(&input); err != nil {
			log.Panicln(err)
		}
	} else {
		if err := c.BindJSON(&input); err != nil {
			log.Panicln(err)
		}
	}

	for i, a := range models.DB {

		if a.ID == id {

			if input.Title != "" {
				models.DB[i].Title = input.Title
			}

			if input.Author != "" {
				models.DB[i].Author = input.Author
			}

		}
	}

	if c.ContentType() == "application/xml" {
		c.XML(http.StatusOK, gin.H{"data": "Updated successfully "})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Updated successfully "})
	}

}

// DELETE /books/:id
// Delete a book

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, a := range models.DB {
		if a.ID == id {
			models.DB = append(models.DB[:i], models.DB[i+1:]...)
		}
	}

	if c.ContentType() == "application/xml" {
		c.XML(http.StatusOK, gin.H{"data": "deleted successfully "})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "deleted successfully "})
	}

}
