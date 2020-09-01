package book

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// Controller : controller for Books
type Controller struct {
	BookService Service
}

// ProvideBookController : provides ProvideBookController
func ProvideBookController(s Service) Controller {
	return Controller{BookService: s}
}

// FindBookByID : find by book id
func (c *Controller) FindBookByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	fmt.Println("fetching book with %id ", id)
	var data Book = c.BookService.FindBookByID(uint(id))
	ctx.JSON(http.StatusOK, data)
}

// FindBookByTitle : find book by title
func (c *Controller) FindBookByTitle(ctx *gin.Context) {
	var data Book = c.BookService.FindBookByTitle(ctx.Param("title"))
	ctx.JSON(http.StatusOK, data)
}

// FindBooks : gets all books
func (c *Controller) FindBooks(ctx *gin.Context) {
	log.Println("getting all books")
	var data []Book = c.BookService.GetAllBooks()
	log.Println(data)
	ctx.JSON(http.StatusOK, data)
}
