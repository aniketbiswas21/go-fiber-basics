package book

import (
	"github.com/aniketbiswas21/go-fiber-basics/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
	c.Send("A single book")
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	// var book Book
	// book.Title = "1984"
	// book.Author = "George Orwell"
	// book.Rating = 5

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(500).Send(err)
		return
	}

	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book found with given ID")
		return
	}
	db.Delete(&book)
	c.Send("Book successfully deleted")
}
