package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

var books = []book{
	{Id: "1", Title: "The Wizard of the Crow", Author: "Ngugi wa Thiong'o", Quantity: 3},
	{Id: "2", Title: "The Adventures of Thiga", Author: "C.M. Muriithi", Quantity: 5},
	{Id: "3", Title: "Kaka Sungura na Wenzake", Author: "David N. Michuki", Quantity: 4},
}

func getBooks(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context)  {
	var newBook book

	// err := c.BindJSON(&newBook)
	if err := c.BindJSON(&newBook); err!=nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated,  newBook)
}

func getBookById(c *gin.Context)  {
	id := c.Param("id")
	book, err := bookById(id)
	if err!=nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func bookById(id string) (*book, error) {
	for i, b := range books {
		if b.Id==id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func borrowBook(c *gin.Context)  {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing a query param"})
		return
	}

	book, err := bookById(id)
	if err!=nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity<=0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK,  book)
}

func returnBook(c *gin.Context)  {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing a query param"})
		return
	}

	book, err := bookById(id)
	if err!=nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found. try creating the book instead"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func main()  {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", createBook)
	router.PATCH("/borrow", borrowBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:5000")
}